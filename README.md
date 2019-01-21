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
```# metrics
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
Markdown                         1            186              0            679
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            306             94           1031
-------------------------------------------------------------------------------
```# metrics
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
Markdown                         1            279              0           1024
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            399             94           1376
-------------------------------------------------------------------------------
```# metrics
--
    import "github.com/reddotpay/metrics"


## Usage

```go
var (

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
Markdown                         1            372              0           1367
Go                               2             85             81            203
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            492             94           1718
-------------------------------------------------------------------------------
```# metrics
--
    import "github.com/reddotpay/metrics"


## Usage

```go
var (

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
Markdown                         1            465              0           1710
Go                               2             85             81            203
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            585             94           2061
-------------------------------------------------------------------------------
```# metrics
--
    import "github.com/reddotpay/metrics"


## Usage

```go
var (

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
Markdown                         1            558              0           2053
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            678             94           2405
-------------------------------------------------------------------------------
```# metrics
--
    import "github.com/reddotpay/metrics"


## Usage

```go
var (

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
Markdown                         1            651              0           2396
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            771             94           2748
-------------------------------------------------------------------------------
```# metrics
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
Markdown                         1            744              0           2741
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            864             94           3093
-------------------------------------------------------------------------------
```# metrics
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
Markdown                         1            837              0           3086
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9            957             94           3438
-------------------------------------------------------------------------------
```# metrics
--
    import "github.com/reddotpay/metrics"


## Usage

```go
var (

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
Markdown                         1            930              0           3429
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9           1050             94           3781
-------------------------------------------------------------------------------
```# metrics
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
Markdown                         1           1023              0           3774
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9           1143             94           4126
-------------------------------------------------------------------------------
```# metrics
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
Markdown                         1           1116              0           4119
Go                               2             85             81            204
Plain Text                       1              0              0             71
BASH                             4             20             13             45
Makefile                         1             15              0             32
-------------------------------------------------------------------------------
TOTAL                            9           1236             94           4471
-------------------------------------------------------------------------------
```