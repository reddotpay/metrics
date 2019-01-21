package metrics_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"

	"github.com/reddotpay/metrics"
	"github.com/stretchr/testify/assert"
)

var m metrics.Metrics

func init() {
	m = metrics.New("appname", "method", "path", "traceID", "apiKey")
}

func TestMetrics_New(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("appname", m.Value["appname"])
	assert.Equal("method", m.Value["method"])
	assert.Equal("path", m.Value["path"])
	assert.Equal("traceID", m.Value["traceID"])
	assert.Equal("c9ff119073ea2567730fb42e3a4fe805", m.Value["apiKey"])
}

func TestMetrics_SetAWSResources(t *testing.T) {
	m.SetAWSResources(metrics.AWSResourceAPIGateway)
	assert := assert.New(t)
	_, ok := m.Value[string(metrics.AWSResourceAPIGateway)]
	assert.True(ok)
	_, ok = m.Value[string(metrics.AWSResourceLambda)]
	assert.False(ok)
}

func TestMetrics_SetDynamoDBReadUsage(t *testing.T) {
	m.SetDynamoDBReadUsage()
	assert := assert.New(t)
	assert.Equal("1", m.Value["dynamodbRead"])
}

func TestMetrics_SetDynamoDBWriteUsage(t *testing.T) {
	m.SetDynamoDBWriteUsage()
	m.SetDynamoDBWriteUsage()
	assert := assert.New(t)
	assert.Equal("2", m.Value["dynamodbWrite"])
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
	metrics.Firehose = metrics.NewClient()
	assert := assert.New(t)
	m.Send(context.Background())
	assert.NotEmpty(m.Value["duration"])
}
