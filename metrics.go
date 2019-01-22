package metrics

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

var (
	// Firehose contains firehose client
	Firehose *firehose.Firehose

	// FirehoseStreamName contains firehose stream name
	FirehoseStreamName string
)

// Metrics defines metrics to send to firehose
type Metrics struct {
	dynamodbRead  float64
	dynamodbWrite float64

	Value map[string]string
}

// New initialises a new `Metrics`
func New(appname, method, path, traceID, apiKey string) Metrics {
	md5 := md5.New()
	md5.Write([]byte(apiKey))

	re := regexp.MustCompile(`.*(Root=\w+-\w+-\w+).*`)
	match := re.FindStringSubmatch(traceID)
	return Metrics{
		dynamodbRead:  0,
		dynamodbWrite: 0,
		Value: map[string]string{
			"appname": appname,
			"method":  method,
			"path":    path,
			"traceID": match[1],
			"apiKey":  hex.EncodeToString(md5.Sum(nil)),
		},
	}
}

// SetAWSResources sets aws resource
func (metrics Metrics) SetAWSResources(resources ...string) {
	for _, resource := range resources {
		var found bool
		for i := range awsResources {
			if awsResources[i] == resource {
				metrics.Value[resource] = ""
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
	metrics.Value["dynamodbRead"] = fmt.Sprintf("%.2f", metrics.dynamodbRead)
}

// SetDynamoDBWriteUsage sets dynamodb write usage
func (metrics *Metrics) SetDynamoDBWriteUsage(usage float64) {
	metrics.dynamodbWrite = metrics.dynamodbWrite + usage
	metrics.Value["dynamodbWrite"] = fmt.Sprintf("%.2f", metrics.dynamodbWrite)
}

// SetDuration sets lambda duration
func (metrics *Metrics) SetDuration(start time.Time) {
	metrics.Value["duration"] = fmt.Sprintf("%.2f", float64(time.Since(start).Nanoseconds())/1000000)
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
	b, _ := json.Marshal(metrics.Value)
	_, err := Firehose.PutRecordWithContext(ctx, &firehose.PutRecordInput{
		DeliveryStreamName: aws.String(FirehoseStreamName),
		Record: &firehose.Record{
			Data: b,
		},
	})

	return err
}
