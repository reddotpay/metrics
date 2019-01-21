package metrics

var awsResources = map[string]int{
	"Alexa":                     1,
	"Amazon MQ":                 1,
	"API Gateway":               1,
	"Application Auto Scaling":  1,
	"AppStream 2.0":             1,
	"AppSync":                   1,
	"Athena":                    1,
	"Auto Scaling":              1,
	"AWS Batch":                 1,
	"Budgets":                   1,
	"Certificate Manager":       1,
	"AWS Cloud9":                1,
	"CloudFormation":            1,
	"CloudFront":                1,
	"AWS Cloud Map":             1,
	"CloudTrail":                1,
	"CloudWatch":                1,
	"CodeBuild":                 1,
	"CodeCommit":                1,
	"CodeDeploy":                1,
	"CodePipeline":              1,
	"Amazon Cognito":            1,
	"AWS Config":                1,
	"Data Pipeline":             1,
	"DAX":                       1,
	"Directory Service":         1,
	"Data Lifecycle Manager":    1,
	"DMS":                       1,
	"DocumentDB":                1,
	"DynamoDB":                  1,
	"EC2":                       1,
	"ECR":                       1,
	"ECS":                       1,
	"EFS":                       1,
	"EKS":                       1,
	"ElastiCache":               1,
	"Elasticsearch":             1,
	"Elastic Beanstalk":         1,
	"Elastic Load Balancing":    1,
	"Elastic Load Balancing V2": 1,
	"EMR":                       1,
	"Amazon GameLift":           1,
	"Glue":                      1,
	"GuardDuty":                 1,
	"IAM":                       1,
	"Amazon Inspector":          1,
	"AWS IoT":                   1,
	"AWS IoT 1-Click":           1,
	"AWS IoT Analytics":         1,
	"Kinesis":                   1,
	"KMS":                       1,
	"Lambda":                    1,
	"Neptune":                   1,
	"OpsWorks":                  1,
	"RDS":                       1,
	"Amazon Redshift":           1,
	"Route 53":                  1,
	"S3":                        1,
	"SageMaker":                 1,
	"Secrets Manager":           1,
	"Service Catalog":           1,
	"SES":                       1,
	"SimpleDB":                  1,
	"SNS":                       1,
	"SQS":                       1,
	"Step Functions":            1,
	"Systems Manager":           1,
	"WAF":                       1,
	"WAF Regional":              1,
	"WorkSpaces":                1,
	"Shared Property Types":     1,
}

// AWSResource defines aws resources
type AWSResource string

// Alexa defines Alexa resource
const Alexa AWSResource = "Alexa"

// AmazonMQ defines Amazon MQ resource
const AmazonMQ AWSResource = "Amazon MQ"

// APIGateway defines API Gateway resource
const APIGateway AWSResource = "API Gateway"

// ApplicationAutoScaling defines Application Auto Scaling resource
const ApplicationAutoScaling AWSResource = "Application Auto Scaling"

// AppStream20 defines AppStream 2.0 resource
const AppStream20 AWSResource = "AppStream 2.0"

// AppSync defines AppSync resource
const AppSync AWSResource = "AppSync"

// Athena defines Athena resource
const Athena AWSResource = "Athena"

// AutoScaling defines Auto Scaling resource
const AutoScaling AWSResource = "Auto Scaling"

// AWSBatch defines AWS Batch resource
const AWSBatch AWSResource = "AWS Batch"

// Budgets defines Budgets resource
const Budgets AWSResource = "Budgets"

// CertificateManager defines Certificate Manager resource
const CertificateManager AWSResource = "Certificate Manager"

// AWSCloud9 defines AWS Cloud9 resource
const AWSCloud9 AWSResource = "AWS Cloud9"

// CloudFormation defines CloudFormation resource
const CloudFormation AWSResource = "CloudFormation"

// CloudFront defines CloudFront resource
const CloudFront AWSResource = "CloudFront"

// AWSCloudMap defines AWS Cloud Map resource
const AWSCloudMap AWSResource = "AWS Cloud Map"

// CloudTrail defines CloudTrail resource
const CloudTrail AWSResource = "CloudTrail"

// CloudWatch defines CloudWatch resource
const CloudWatch AWSResource = "CloudWatch"

// CodeBuild defines CodeBuild resource
const CodeBuild AWSResource = "CodeBuild"

// CodeCommit defines CodeCommit resource
const CodeCommit AWSResource = "CodeCommit"

// CodeDeploy defines CodeDeploy resource
const CodeDeploy AWSResource = "CodeDeploy"

// CodePipeline defines CodePipeline resource
const CodePipeline AWSResource = "CodePipeline"

// AmazonCognito defines Amazon Cognito resource
const AmazonCognito AWSResource = "Amazon Cognito"

// AWSConfig defines AWS Config resource
const AWSConfig AWSResource = "AWS Config"

// DataPipeline defines Data Pipeline resource
const DataPipeline AWSResource = "Data Pipeline"

// DAX defines DAX resource
const DAX AWSResource = "DAX"

// DirectoryService defines Directory Service resource
const DirectoryService AWSResource = "Directory Service"

// DataLifecycleManager defines Data Lifecycle Manager resource
const DataLifecycleManager AWSResource = "Data Lifecycle Manager"

// DMS defines DMS resource
const DMS AWSResource = "DMS"

// DocumentDB defines DocumentDB resource
const DocumentDB AWSResource = "DocumentDB"

// DynamoDB defines DynamoDB resource
const DynamoDB AWSResource = "DynamoDB"

// EC2 defines EC2 resource
const EC2 AWSResource = "EC2"

// ECR defines ECR resource
const ECR AWSResource = "ECR"

// ECS defines ECS resource
const ECS AWSResource = "ECS"

// EFS defines EFS resource
const EFS AWSResource = "EFS"

// EKS defines EKS resource
const EKS AWSResource = "EKS"

// ElastiCache defines ElastiCache resource
const ElastiCache AWSResource = "ElastiCache"

// Elasticsearch defines Elasticsearch resource
const Elasticsearch AWSResource = "Elasticsearch"

// ElasticBeanstalk defines Elastic Beanstalk resource
const ElasticBeanstalk AWSResource = "Elastic Beanstalk"

// ElasticLoadBalancing defines Elastic Load Balancing resource
const ElasticLoadBalancing AWSResource = "Elastic Load Balancing"

// ElasticLoadBalancingV2 defines Elastic Load Balancing V2 resource
const ElasticLoadBalancingV2 AWSResource = "Elastic Load Balancing V2"

// EMR defines EMR resource
const EMR AWSResource = "EMR"

// AmazonGameLift defines Amazon GameLift resource
const AmazonGameLift AWSResource = "Amazon GameLift"

// Glue defines Glue resource
const Glue AWSResource = "Glue"

// GuardDuty defines GuardDuty resource
const GuardDuty AWSResource = "GuardDuty"

// IAM defines IAM resource
const IAM AWSResource = "IAM"

// AmazonInspector defines Amazon Inspector resource
const AmazonInspector AWSResource = "Amazon Inspector"

// AWSIoT defines AWS IoT resource
const AWSIoT AWSResource = "AWS IoT"

// AWSIoT1Click defines AWS IoT 1-Click resource
const AWSIoT1Click AWSResource = "AWS IoT 1-Click"

// AWSIoTAnalytics defines AWS IoT Analytics resource
const AWSIoTAnalytics AWSResource = "AWS IoT Analytics"

// Kinesis defines Kinesis resource
const Kinesis AWSResource = "Kinesis"

// KMS defines KMS resource
const KMS AWSResource = "KMS"

// Lambda defines Lambda resource
const Lambda AWSResource = "Lambda"

// Neptune defines Neptune resource
const Neptune AWSResource = "Neptune"

// OpsWorks defines OpsWorks resource
const OpsWorks AWSResource = "OpsWorks"

// RDS defines RDS resource
const RDS AWSResource = "RDS"

// AmazonRedshift defines Amazon Redshift resource
const AmazonRedshift AWSResource = "Amazon Redshift"

// Route53 defines Route 53 resource
const Route53 AWSResource = "Route 53"

// S3 defines S3 resource
const S3 AWSResource = "S3"

// SageMaker defines SageMaker resource
const SageMaker AWSResource = "SageMaker"

// SecretsManager defines Secrets Manager resource
const SecretsManager AWSResource = "Secrets Manager"

// ServiceCatalog defines Service Catalog resource
const ServiceCatalog AWSResource = "Service Catalog"

// SES defines SES resource
const SES AWSResource = "SES"

// SimpleDB defines SimpleDB resource
const SimpleDB AWSResource = "SimpleDB"

// SNS defines SNS resource
const SNS AWSResource = "SNS"

// SQS defines SQS resource
const SQS AWSResource = "SQS"

// StepFunctions defines Step Functions resource
const StepFunctions AWSResource = "Step Functions"

// SystemsManager defines Systems Manager resource
const SystemsManager AWSResource = "Systems Manager"

// WAF defines WAF resource
const WAF AWSResource = "WAF"

// WAFRegional defines WAF Regional resource
const WAFRegional AWSResource = "WAF Regional"

// WorkSpaces defines WorkSpaces resource
const WorkSpaces AWSResource = "WorkSpaces"

// SharedPropertyTypes defines Shared Property Types resource
const SharedPropertyTypes AWSResource = "Shared Property Types"
