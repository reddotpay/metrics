package metrics_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"

	"github.com/reddotpay/metrics"
	"github.com/stretchr/testify/assert"
)

var m metrics.Metrics

func init() {
	m = metrics.New("appname", "method", "path", "requestId", "Root=1-5c4671a9-2a7f8aa8bb44bba86ddc3550;Parent=a10d99b18d2691df;Sampled=1", "apiKey")
}

func TestMetrics_New(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("appname", m.Value["appname"])
	assert.Equal("method", m.Value["method"])
	assert.Equal("path", m.Value["path"])
	assert.Equal("requestId", m.Value["request_id"])
	assert.Equal("Root=1-5c4671a9-2a7f8aa8bb44bba86ddc3550", m.Value["trace_id"])
	assert.Equal("c9ff119073ea2567730fb42e3a4fe805", m.Value["api_key"])
}

func TestMetrics_SetAWSResources(t *testing.T) {
	m.SetAWSResources("API Gateway", "Lambda")
	assert := assert.New(t)
	assert.Equal([]string{"API Gateway", "Lambda"}, m.Resources)
}

func TestMetrics_SetDynamoDBReadUsage(t *testing.T) {
	m.SetDynamoDBReadUsage(1)
	assert := assert.New(t)
	assert.Equal(1.0, m.Value["dynamodb.read.read"])
}

func TestMetrics_SetDynamoDBWriteUsage(t *testing.T) {
	m.SetDynamoDBWriteUsage(0.5)
	assert := assert.New(t)
	assert.Equal(0.5, m.Value["dynamodb.write.write"])
}

func TestMetrics_SetDuration(t *testing.T) {
	m.SetDuration(time.Now().Add(-1 * time.Second))
	assert := assert.New(t)
	assert.Equal(1000.00, m.Value["lambda.duration.ms"])
}

func TestMetrics_SetLambdaMemorySize(t *testing.T) {
	m.SetLambdaMemorySize("128")
	assert := assert.New(t)
	assert.Equal(float64(128.00), m.Value["lambda.memory.MB"])
}

func TestMetrics_NewClient(t *testing.T) {
	client := metrics.NewClient()
	assert := assert.New(t)
	assert.IsType(&firehose.Firehose{}, client)
}

func TestMetrics_NewClientWithConfig(t *testing.T) {
	client := metrics.NewClientWithConfig(aws.Config{})
	assert := assert.New(t)
	assert.IsType(&firehose.Firehose{}, client)
}

func TestMetrics_Send(t *testing.T) {
	metrics.Firehose = &mockFirehose{}
	m = metrics.New("appname", "method", "path", "requestId", "Root=1-5c4671a9-2a7f8aa8bb44bba86ddc3550;Parent=a10d99b18d2691df;Sampled=1", "apiKey")
	m.SetAWSResources("API Gateway", "Lambda", "DynamoDB")
	m.SetLambdaMemorySize("128")
	m.SetDuration(time.Now().Add(-1 * time.Second))
	m.SetDynamoDBWriteUsage(2)
	err := m.Send(context.Background())
	assert := assert.New(t)
	assert.NoError(err)
}

type mockFirehose struct {
	firehoseiface.FirehoseAPI
}

func (fh *mockFirehose) PutRecordBatchWithContext(ctx aws.Context, input *firehose.PutRecordBatchInput, req ...request.Option) (*firehose.PutRecordBatchOutput, error) {
	return &firehose.PutRecordBatchOutput{}, nil
}
