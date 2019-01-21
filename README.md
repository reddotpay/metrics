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
func NewClient(streamName string) *firehose.Firehose
```
NewClient creates a new firehose client with default config

#### func  NewClientWithConfig

```go
func NewClientWithConfig(streamName string, config aws.Config) *firehose.Firehose
```
NewClientWithConfig creates a new firehose client with given config

#### type AWSResource

```go
type AWSResource string
```

AWSResource defines aws resources

```go
const APIGateway AWSResource = "API Gateway"
```
APIGateway defines API Gateway resource

```go
const AWSBatch AWSResource = "AWS Batch"
```
AWSBatch defines AWS Batch resource

```go
const AWSCloud9 AWSResource = "AWS Cloud9"
```
AWSCloud9 defines AWS Cloud9 resource

```go
const AWSCloudMap AWSResource = "AWS Cloud Map"
```
AWSCloudMap defines AWS Cloud Map resource

```go
const AWSConfig AWSResource = "AWS Config"
```
AWSConfig defines AWS Config resource

```go
const AWSIoT AWSResource = "AWS IoT"
```
AWSIoT defines AWS IoT resource

```go
const AWSIoT1Click AWSResource = "AWS IoT 1-Click"
```
AWSIoT1Click defines AWS IoT 1-Click resource

```go
const AWSIoTAnalytics AWSResource = "AWS IoT Analytics"
```
AWSIoTAnalytics defines AWS IoT Analytics resource

```go
const Alexa AWSResource = "Alexa"
```
Alexa defines Alexa resource

```go
const AmazonCognito AWSResource = "Amazon Cognito"
```
AmazonCognito defines Amazon Cognito resource

```go
const AmazonGameLift AWSResource = "Amazon GameLift"
```
AmazonGameLift defines Amazon GameLift resource

```go
const AmazonInspector AWSResource = "Amazon Inspector"
```
AmazonInspector defines Amazon Inspector resource

```go
const AmazonMQ AWSResource = "Amazon MQ"
```
AmazonMQ defines Amazon MQ resource

```go
const AmazonRedshift AWSResource = "Amazon Redshift"
```
AmazonRedshift defines Amazon Redshift resource

```go
const AppStream20 AWSResource = "AppStream 2.0"
```
AppStream20 defines AppStream 2.0 resource

```go
const AppSync AWSResource = "AppSync"
```
AppSync defines AppSync resource

```go
const ApplicationAutoScaling AWSResource = "Application Auto Scaling"
```
ApplicationAutoScaling defines Application Auto Scaling resource

```go
const Athena AWSResource = "Athena"
```
Athena defines Athena resource

```go
const AutoScaling AWSResource = "Auto Scaling"
```
AutoScaling defines Auto Scaling resource

```go
const Budgets AWSResource = "Budgets"
```
Budgets defines Budgets resource

```go
const CertificateManager AWSResource = "Certificate Manager"
```
CertificateManager defines Certificate Manager resource

```go
const CloudFormation AWSResource = "CloudFormation"
```
CloudFormation defines CloudFormation resource

```go
const CloudFront AWSResource = "CloudFront"
```
CloudFront defines CloudFront resource

```go
const CloudTrail AWSResource = "CloudTrail"
```
CloudTrail defines CloudTrail resource

```go
const CloudWatch AWSResource = "CloudWatch"
```
CloudWatch defines CloudWatch resource

```go
const CodeBuild AWSResource = "CodeBuild"
```
CodeBuild defines CodeBuild resource

```go
const CodeCommit AWSResource = "CodeCommit"
```
CodeCommit defines CodeCommit resource

```go
const CodeDeploy AWSResource = "CodeDeploy"
```
CodeDeploy defines CodeDeploy resource

```go
const CodePipeline AWSResource = "CodePipeline"
```
CodePipeline defines CodePipeline resource

```go
const DAX AWSResource = "DAX"
```
DAX defines DAX resource

```go
const DMS AWSResource = "DMS"
```
DMS defines DMS resource

```go
const DataLifecycleManager AWSResource = "Data Lifecycle Manager"
```
DataLifecycleManager defines Data Lifecycle Manager resource

```go
const DataPipeline AWSResource = "Data Pipeline"
```
DataPipeline defines Data Pipeline resource

```go
const DirectoryService AWSResource = "Directory Service"
```
DirectoryService defines Directory Service resource

```go
const DocumentDB AWSResource = "DocumentDB"
```
DocumentDB defines DocumentDB resource

```go
const DynamoDB AWSResource = "DynamoDB"
```
DynamoDB defines DynamoDB resource

```go
const EC2 AWSResource = "EC2"
```
EC2 defines EC2 resource

```go
const ECR AWSResource = "ECR"
```
ECR defines ECR resource

```go
const ECS AWSResource = "ECS"
```
ECS defines ECS resource

```go
const EFS AWSResource = "EFS"
```
EFS defines EFS resource

```go
const EKS AWSResource = "EKS"
```
EKS defines EKS resource

```go
const EMR AWSResource = "EMR"
```
EMR defines EMR resource

```go
const ElastiCache AWSResource = "ElastiCache"
```
ElastiCache defines ElastiCache resource

```go
const ElasticBeanstalk AWSResource = "Elastic Beanstalk"
```
ElasticBeanstalk defines Elastic Beanstalk resource

```go
const ElasticLoadBalancing AWSResource = "Elastic Load Balancing"
```
ElasticLoadBalancing defines Elastic Load Balancing resource

```go
const ElasticLoadBalancingV2 AWSResource = "Elastic Load Balancing V2"
```
ElasticLoadBalancingV2 defines Elastic Load Balancing V2 resource

```go
const Elasticsearch AWSResource = "Elasticsearch"
```
Elasticsearch defines Elasticsearch resource

```go
const Glue AWSResource = "Glue"
```
Glue defines Glue resource

```go
const GuardDuty AWSResource = "GuardDuty"
```
GuardDuty defines GuardDuty resource

```go
const IAM AWSResource = "IAM"
```
IAM defines IAM resource

```go
const KMS AWSResource = "KMS"
```
KMS defines KMS resource

```go
const Kinesis AWSResource = "Kinesis"
```
Kinesis defines Kinesis resource

```go
const Lambda AWSResource = "Lambda"
```
Lambda defines Lambda resource

```go
const Neptune AWSResource = "Neptune"
```
Neptune defines Neptune resource

```go
const OpsWorks AWSResource = "OpsWorks"
```
OpsWorks defines OpsWorks resource

```go
const RDS AWSResource = "RDS"
```
RDS defines RDS resource

```go
const Route53 AWSResource = "Route 53"
```
Route53 defines Route 53 resource

```go
const S3 AWSResource = "S3"
```
S3 defines S3 resource

```go
const SES AWSResource = "SES"
```
SES defines SES resource

```go
const SNS AWSResource = "SNS"
```
SNS defines SNS resource

```go
const SQS AWSResource = "SQS"
```
SQS defines SQS resource

```go
const SageMaker AWSResource = "SageMaker"
```
SageMaker defines SageMaker resource

```go
const SecretsManager AWSResource = "Secrets Manager"
```
SecretsManager defines Secrets Manager resource

```go
const ServiceCatalog AWSResource = "Service Catalog"
```
ServiceCatalog defines Service Catalog resource

```go
const SharedPropertyTypes AWSResource = "Shared Property Types"
```
SharedPropertyTypes defines Shared Property Types resource

```go
const SimpleDB AWSResource = "SimpleDB"
```
SimpleDB defines SimpleDB resource

```go
const StepFunctions AWSResource = "Step Functions"
```
StepFunctions defines Step Functions resource

```go
const SystemsManager AWSResource = "Systems Manager"
```
SystemsManager defines Systems Manager resource

```go
const WAF AWSResource = "WAF"
```
WAF defines WAF resource

```go
const WAFRegional AWSResource = "WAF Regional"
```
WAFRegional defines WAF Regional resource

```go
const WorkSpaces AWSResource = "WorkSpaces"
```
WorkSpaces defines WorkSpaces resource

#### type Metrics

```go
type Metrics struct {
	Value map[string]string
}
```

Metrics defines metrics to send to firehose

#### func  New

```go
func New(method, path, traceID, apiKey string) Metrics
```
New initialises a new `Metrics`

#### func (Metrics) Send

```go
func (metrics Metrics) Send(ctx context.Context) error
```
Send sends `Metrics` to firehose stream

#### func (Metrics) SetAWSResources

```go
func (metrics Metrics) SetAWSResources(resources ...AWSResource)
```
SetAWSResources sets aws resource


```
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Markdown                         1             93              0            334
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            213             94            686
-------------------------------------------------------------------------------
```