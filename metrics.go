package metrics

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
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
	start time.Time
	Value map[string]string
}

// New initialises a new `Metrics`
func New(method, path, traceID, apiKey string) Metrics {
	md5 := md5.New()
	md5.Write([]byte(apiKey))
	return Metrics{
		start: time.Now(),
		Value: map[string]string{
			"method":  method,
			"path":    path,
			"traceID": traceID,
			"apiKey":  hex.EncodeToString(md5.Sum(nil)),
		},
	}
}

// SetAWSResources sets aws resource
func (metrics Metrics) SetAWSResources(resources ...AWSResource) {
	for _, resource := range resources {
		if _, ok := awsResources[string(resource)]; !ok {
			panic("invalid AWS resource")
		}
		metrics.Value[string(resource)] = ""
	}
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
func (metrics Metrics) Send(ctx context.Context) error {
	// Set duration in milliseconds
	metrics.Value["duration"] = strconv.FormatInt(time.Since(metrics.start).Nanoseconds()/1000000, 10)

	b, _ := json.Marshal(metrics.Value)
	_, err := Firehose.PutRecordWithContext(ctx, &firehose.PutRecordInput{
		DeliveryStreamName: aws.String(FirehoseStreamName),
		Record: &firehose.Record{
			Data: b,
		},
	})

	return err
}
