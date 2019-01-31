# metrics
--
    import "github.com/reddotpay/metrics"


## Usage

```go
var (
	// Firehose contains firehose client
	Firehose *firehose.Firehose

	// FirehoseStreamName contains firehose stream name
	FirehoseStreamName string
)
```

#### func  NewClient

```go
func NewClient() *firehose.Firehose
```
NewClient creates a new firehose client with default config

#### func  NewClientWithConfig

```go
func NewClientWithConfig(config aws.Config) *firehose.Firehose
```
NewClientWithConfig creates a new firehose client with given config

#### type Metrics

```go
type Metrics struct {
	Value map[string]string
}
```

Metrics defines metrics to send to firehose

#### func  New

```go
func New(appname, method, path, traceID, apiKey string) Metrics
```
New initialises a new `Metrics`

#### func (*Metrics) Send

```go
func (metrics *Metrics) Send(ctx context.Context) error
```
Send sends `Metrics` to firehose stream

#### func (Metrics) SetAWSResources

```go
func (metrics Metrics) SetAWSResources(resources ...string)
```
SetAWSResources sets aws resource

#### func (*Metrics) SetDuration

```go
func (metrics *Metrics) SetDuration(start time.Time)
```
SetDuration sets lambda duration

#### func (*Metrics) SetDynamoDBReadUsage

```go
func (metrics *Metrics) SetDynamoDBReadUsage(usage float64)
```
SetDynamoDBReadUsage sets dynamodb read usage

#### func (*Metrics) SetDynamoDBWriteUsage

```go
func (metrics *Metrics) SetDynamoDBWriteUsage(usage float64)
```
SetDynamoDBWriteUsage sets dynamodb write usage

#### func (*Metrics) SetLambdaMemorySize

```go
func (metrics *Metrics) SetLambdaMemorySize(memory string)
```
SetLambdaMemorySize sets lambda duration


```
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             32             12            219
Plain Text                       1              0              0             71
Markdown                         1             27              0             65
BASH                             4             17             12             33
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                           10             91             24            420
-------------------------------------------------------------------------------
```