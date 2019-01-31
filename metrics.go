package metrics

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
)

var (
	// Firehose contains firehose client
	Firehose firehoseiface.FirehoseAPI

	// FirehoseStreamName contains firehose stream name
	FirehoseStreamName string
)

// Metrics defines metrics to send to firehose
type Metrics struct {
	dynamodbRead  float64
	dynamodbWrite float64

	Value     map[string]interface{}
	Resources []string
}

// New initialises a new `Metrics`
func New(appname, method, path, requestID, traceID, apiKey string) Metrics {
	md5 := md5.New()
	md5.Write([]byte(apiKey))

	re := regexp.MustCompile(`.*(Root=\w+-\w+-\w+).*`)
	match := re.FindStringSubmatch(traceID)
	return Metrics{
		dynamodbRead:  0,
		dynamodbWrite: 0,
		Value: map[string]interface{}{
			"appname":    appname,
			"method":     method,
			"path":       path,
			"request_id": requestID,
			"trace_id":   match[1],
			"api_key":    hex.EncodeToString(md5.Sum(nil)),
		},
		Resources: []string{},
	}
}

// SetAWSResources sets aws resource
func (metrics *Metrics) SetAWSResources(resources ...string) {
	for _, resource := range resources {
		var found bool
		for i := range awsResources {
			if awsResources[i] == resource {
				metrics.Resources = append(metrics.Resources, resource)
				found = true
				break
			}
		}

		if !found {
			panic("invalid AWS resource")
		}
	}
}

// SetDynamoDBReadUsage sets dynamodb read usage
func (metrics *Metrics) SetDynamoDBReadUsage(usage float64) {
	metrics.dynamodbRead = metrics.dynamodbRead + usage
	metrics.Value["dynamodb.read.read"] = fmt.Sprintf("%.2f", metrics.dynamodbRead)
}

// SetDynamoDBWriteUsage sets dynamodb write usage
func (metrics *Metrics) SetDynamoDBWriteUsage(usage float64) {
	metrics.dynamodbWrite = metrics.dynamodbWrite + usage
	metrics.Value["dynamodb.write.write"] = fmt.Sprintf("%.2f", metrics.dynamodbWrite)
}

// SetDuration sets lambda duration
func (metrics *Metrics) SetDuration(start time.Time) {
	metrics.Value["lambda.duration.ms"] = fmt.Sprintf("%.2f", float64(time.Since(start).Nanoseconds())/1000000)
}

// SetLambdaMemorySize sets lambda duration
func (metrics *Metrics) SetLambdaMemorySize(memory string) {
	f, _ := strconv.ParseFloat(memory, 64)
	metrics.Value["lambda.memory.MB"] = f
}

// NewClient creates a new firehose client with default config
func NewClient() *firehose.Firehose {
	return firehose.New(session.New(&aws.Config{}))
}

// NewClientWithConfig creates a new firehose client with given config
func NewClientWithConfig(config aws.Config) *firehose.Firehose {
	return firehose.New(session.New(&config))
}

// Send sends `Metrics` to firehose stream
func (metrics *Metrics) Send(ctx context.Context) error {
	records := []*firehose.Record{}
	for i := range metrics.Resources {
		resource := formatResource(metrics.Resources[i])

		var hasUsage bool
		re := regexp.MustCompile(`^` + resource + `\.(\w+)\.(\w+)$`)
		for k, v := range metrics.Value {
			if re.MatchString(k) {
				hasUsage = true

				ss := strings.Split(k, ".")
				b, _ := json.Marshal(map[string]interface{}{
					"appname":    metrics.Value["appname"],
					"method":     metrics.Value["method"],
					"path":       metrics.Value["path"],
					"request_id": metrics.Value["request_id"],
					"trace_id":   metrics.Value["trace_id"],
					"api_key":    metrics.Value["api_key"],
					"resource":   resource,
					"usage":      v,
					"usage_unit": ss[2],
					"metric":     ss[1],
				})
				records = append(records, &firehose.Record{
					Data: b,
				})
			}
		}
		if !hasUsage {
			b, _ := json.Marshal(map[string]interface{}{
				"appname":    metrics.Value["appname"],
				"method":     metrics.Value["method"],
				"path":       metrics.Value["path"],
				"request_id": metrics.Value["request_id"],
				"trace_id":   metrics.Value["trace_id"],
				"api_key":    metrics.Value["api_key"],
				"resource":   resource,
				"usage":      1,
				"usage_unit": "call",
				"metric":     "call",
			})
			records = append(records, &firehose.Record{
				Data: b,
			})
		}
	}

	_, err := Firehose.PutRecordBatchWithContext(ctx, &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(FirehoseStreamName),
		Records:            records,
	})
	return err
}

// formatResource formats the resource string to lowercase and remove spaces
func formatResource(resource string) string {
	return strings.Replace(strings.ToLower(resource), " ", "_", -1)
}
