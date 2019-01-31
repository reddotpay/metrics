package metrics_test

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"

	"github.com/reddotpay/metrics"
	"github.com/stretchr/testify/assert"
)

var m metrics.Metrics

func init() {
	m = metrics.New("appname", "method", "path", "Root=1-5c4671a9-2a7f8aa8bb44bba86ddc3550;Parent=a10d99b18d2691df;Sampled=1", "apiKey")
}

func TestMetrics_New(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("appname", m.Value["appname"])
	assert.Equal("method", m.Value["method"])
	assert.Equal("path", m.Value["path"])
	assert.Equal("Root=1-5c4671a9-2a7f8aa8bb44bba86ddc3550", m.Value["traceID"])
	assert.Equal("c9ff119073ea2567730fb42e3a4fe805", m.Value["apiKey"])
}

func TestMetrics_SetAWSResources(t *testing.T) {
	m.SetAWSResources("API Gateway", "Lambda")
	assert := assert.New(t)
	_, ok := m.Value["API Gateway"]
	assert.True(ok)
	_, ok = m.Value["Lambda"]
	assert.True(ok)
	_, ok = m.Value["DynamoDB"]
	assert.False(ok)
}

func TestMetrics_SetDynamoDBReadUsage(t *testing.T) {
	m.SetDynamoDBReadUsage(1)
	assert := assert.New(t)
	assert.Equal("1.00", m.Value["dynamodbRead"])
}

func TestMetrics_SetDynamoDBWriteUsage(t *testing.T) {
	m.SetDynamoDBWriteUsage(2)
	assert := assert.New(t)
	assert.Equal("2.00", m.Value["dynamodbWrite"])
}

func TestMetrics_SetDuration(t *testing.T) {
	m.SetDuration(time.Now().Add(-1 * time.Second))
	assert := assert.New(t)
	assert.Equal("1000.00", m.Value["duration"])
}

func TestMetrics_SetLambdaMemorySize(t *testing.T) {
	m.SetLambdaMemorySize("128")
	assert := assert.New(t)
	assert.Equal("128", m.Value["memory"])
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
