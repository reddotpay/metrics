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

#### type AWSResource

```go
type AWSResource string
```

AWSResource defines aws resources

```go
const AWSResourceAPIGateway AWSResource = "API Gateway"
```
AWSResourceAPIGateway defines API Gateway resource

```go
const AWSResourceAWSBatch AWSResource = "AWS Batch"
```
AWSResourceAWSBatch defines AWS Batch resource

```go
const AWSResourceAWSCloud9 AWSResource = "AWS Cloud9"
```
AWSResourceAWSCloud9 defines AWS Cloud9 resource

```go
const AWSResourceAWSCloudMap AWSResource = "AWS Cloud Map"
```
AWSResourceAWSCloudMap defines AWS Cloud Map resource

```go
const AWSResourceAWSConfig AWSResource = "AWS Config"
```
AWSResourceAWSConfig defines AWS Config resource

```go
const AWSResourceAWSIoT AWSResource = "AWS IoT"
```
AWSResourceAWSIoT defines AWS IoT resource

```go
const AWSResourceAWSIoT1Click AWSResource = "AWS IoT 1-Click"
```
AWSResourceAWSIoT1Click defines AWS IoT 1-Click resource

```go
const AWSResourceAWSIoTAnalytics AWSResource = "AWS IoT Analytics"
```
AWSResourceAWSIoTAnalytics defines AWS IoT Analytics resource

```go
const AWSResourceAlexa AWSResource = "Alexa"
```
AWSResourceAlexa defines Alexa resource

```go
const AWSResourceAmazonCognito AWSResource = "Amazon Cognito"
```
AWSResourceAmazonCognito defines Amazon Cognito resource

```go
const AWSResourceAmazonGameLift AWSResource = "Amazon GameLift"
```
AWSResourceAmazonGameLift defines Amazon GameLift resource

```go
const AWSResourceAmazonInspector AWSResource = "Amazon Inspector"
```
AWSResourceAmazonInspector defines Amazon Inspector resource

```go
const AWSResourceAmazonMQ AWSResource = "Amazon MQ"
```
AWSResourceAmazonMQ defines Amazon MQ resource

```go
const AWSResourceAmazonRedshift AWSResource = "Amazon Redshift"
```
AWSResourceAmazonRedshift defines Amazon Redshift resource

```go
const AWSResourceAppStream20 AWSResource = "AppStream 2.0"
```
AWSResourceAppStream20 defines AppStream 2.0 resource

```go
const AWSResourceAppSync AWSResource = "AppSync"
```
AWSResourceAppSync defines AppSync resource

```go
const AWSResourceApplicationAutoScaling AWSResource = "Application Auto Scaling"
```
AWSResourceApplicationAutoScaling defines Application Auto Scaling resource

```go
const AWSResourceAthena AWSResource = "Athena"
```
AWSResourceAthena defines Athena resource

```go
const AWSResourceAutoScaling AWSResource = "Auto Scaling"
```
AWSResourceAutoScaling defines Auto Scaling resource

```go
const AWSResourceBudgets AWSResource = "Budgets"
```
AWSResourceBudgets defines Budgets resource

```go
const AWSResourceCertificateManager AWSResource = "Certificate Manager"
```
AWSResourceCertificateManager defines Certificate Manager resource

```go
const AWSResourceCloudFormation AWSResource = "CloudFormation"
```
AWSResourceCloudFormation defines CloudFormation resource

```go
const AWSResourceCloudFront AWSResource = "CloudFront"
```
AWSResourceCloudFront defines CloudFront resource

```go
const AWSResourceCloudTrail AWSResource = "CloudTrail"
```
AWSResourceCloudTrail defines CloudTrail resource

```go
const AWSResourceCloudWatch AWSResource = "CloudWatch"
```
AWSResourceCloudWatch defines CloudWatch resource

```go
const AWSResourceCodeBuild AWSResource = "CodeBuild"
```
AWSResourceCodeBuild defines CodeBuild resource

```go
const AWSResourceCodeCommit AWSResource = "CodeCommit"
```
AWSResourceCodeCommit defines CodeCommit resource

```go
const AWSResourceCodeDeploy AWSResource = "CodeDeploy"
```
AWSResourceCodeDeploy defines CodeDeploy resource

```go
const AWSResourceCodePipeline AWSResource = "CodePipeline"
```
AWSResourceCodePipeline defines CodePipeline resource

```go
const AWSResourceDAX AWSResource = "DAX"
```
AWSResourceDAX defines DAX resource

```go
const AWSResourceDMS AWSResource = "DMS"
```
AWSResourceDMS defines DMS resource

```go
const AWSResourceDataLifecycleManager AWSResource = "Data Lifecycle Manager"
```
AWSResourceDataLifecycleManager defines Data Lifecycle Manager resource

```go
const AWSResourceDataPipeline AWSResource = "Data Pipeline"
```
AWSResourceDataPipeline defines Data Pipeline resource

```go
const AWSResourceDirectoryService AWSResource = "Directory Service"
```
AWSResourceDirectoryService defines Directory Service resource

```go
const AWSResourceDocumentDB AWSResource = "DocumentDB"
```
AWSResourceDocumentDB defines DocumentDB resource

```go
const AWSResourceDynamoDB AWSResource = "DynamoDB"
```
AWSResourceDynamoDB defines DynamoDB resource

```go
const AWSResourceEC2 AWSResource = "EC2"
```
AWSResourceEC2 defines EC2 resource

```go
const AWSResourceECR AWSResource = "ECR"
```
AWSResourceECR defines ECR resource

```go
const AWSResourceECS AWSResource = "ECS"
```
AWSResourceECS defines ECS resource

```go
const AWSResourceEFS AWSResource = "EFS"
```
AWSResourceEFS defines EFS resource

```go
const AWSResourceEKS AWSResource = "EKS"
```
AWSResourceEKS defines EKS resource

```go
const AWSResourceEMR AWSResource = "EMR"
```
AWSResourceEMR defines EMR resource

```go
const AWSResourceElastiCache AWSResource = "ElastiCache"
```
AWSResourceElastiCache defines ElastiCache resource

```go
const AWSResourceElasticBeanstalk AWSResource = "Elastic Beanstalk"
```
AWSResourceElasticBeanstalk defines Elastic Beanstalk resource

```go
const AWSResourceElasticLoadBalancing AWSResource = "Elastic Load Balancing"
```
AWSResourceElasticLoadBalancing defines Elastic Load Balancing resource

```go
const AWSResourceElasticLoadBalancingV2 AWSResource = "Elastic Load Balancing V2"
```
AWSResourceElasticLoadBalancingV2 defines Elastic Load Balancing V2 resource

```go
const AWSResourceElasticsearch AWSResource = "Elasticsearch"
```
AWSResourceElasticsearch defines Elasticsearch resource

```go
const AWSResourceGlue AWSResource = "Glue"
```
AWSResourceGlue defines Glue resource

```go
const AWSResourceGuardDuty AWSResource = "GuardDuty"
```
AWSResourceGuardDuty defines GuardDuty resource

```go
const AWSResourceIAM AWSResource = "IAM"
```
AWSResourceIAM defines IAM resource

```go
const AWSResourceKMS AWSResource = "KMS"
```
AWSResourceKMS defines KMS resource

```go
const AWSResourceKinesis AWSResource = "Kinesis"
```
AWSResourceKinesis defines Kinesis resource

```go
const AWSResourceLambda AWSResource = "Lambda"
```
AWSResourceLambda defines Lambda resource

```go
const AWSResourceNeptune AWSResource = "Neptune"
```
AWSResourceNeptune defines Neptune resource

```go
const AWSResourceOpsWorks AWSResource = "OpsWorks"
```
AWSResourceOpsWorks defines OpsWorks resource

```go
const AWSResourceRDS AWSResource = "RDS"
```
AWSResourceRDS defines RDS resource

```go
const AWSResourceRoute53 AWSResource = "Route 53"
```
AWSResourceRoute53 defines Route 53 resource

```go
const AWSResourceS3 AWSResource = "S3"
```
AWSResourceS3 defines S3 resource

```go
const AWSResourceSES AWSResource = "SES"
```
AWSResourceSES defines SES resource

```go
const AWSResourceSNS AWSResource = "SNS"
```
AWSResourceSNS defines SNS resource

```go
const AWSResourceSQS AWSResource = "SQS"
```
AWSResourceSQS defines SQS resource

```go
const AWSResourceSageMaker AWSResource = "SageMaker"
```
AWSResourceSageMaker defines SageMaker resource

```go
const AWSResourceSecretsManager AWSResource = "Secrets Manager"
```
AWSResourceSecretsManager defines Secrets Manager resource

```go
const AWSResourceServiceCatalog AWSResource = "Service Catalog"
```
AWSResourceServiceCatalog defines Service Catalog resource

```go
const AWSResourceSharedPropertyTypes AWSResource = "Shared Property Types"
```
AWSResourceSharedPropertyTypes defines Shared Property Types resource

```go
const AWSResourceSimpleDB AWSResource = "SimpleDB"
```
AWSResourceSimpleDB defines SimpleDB resource

```go
const AWSResourceStepFunctions AWSResource = "Step Functions"
```
AWSResourceStepFunctions defines Step Functions resource

```go
const AWSResourceSystemsManager AWSResource = "Systems Manager"
```
AWSResourceSystemsManager defines Systems Manager resource

```go
const AWSResourceWAF AWSResource = "WAF"
```
AWSResourceWAF defines WAF resource

```go
const AWSResourceWAFRegional AWSResource = "WAF Regional"
```
AWSResourceWAFRegional defines WAF Regional resource

```go
const AWSResourceWorkSpaces AWSResource = "WorkSpaces"
```
AWSResourceWorkSpaces defines WorkSpaces resource

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
func (metrics Metrics) SetAWSResources(resources ...AWSResource)
```
SetAWSResources sets aws resource

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


```
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Markdown                         1             97              0            344
Go                               3            101             83            270
Plain Text                       1              0              0             71
BASH                             4             21             13             44
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                           10            234             96            761
-------------------------------------------------------------------------------
```