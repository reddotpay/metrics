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

// AWSResourceAlexa defines Alexa resource
const AWSResourceAlexa AWSResource = "Alexa"

// AWSResourceAmazonMQ defines Amazon MQ resource
const AWSResourceAmazonMQ AWSResource = "Amazon MQ"

// AWSResourceAPIGateway defines API Gateway resource
const AWSResourceAPIGateway AWSResource = "API Gateway"

// AWSResourceApplicationAutoScaling defines Application Auto Scaling resource
const AWSResourceApplicationAutoScaling AWSResource = "Application Auto Scaling"

// AWSResourceAppStream20 defines AppStream 2.0 resource
const AWSResourceAppStream20 AWSResource = "AppStream 2.0"

// AWSResourceAppSync defines AppSync resource
const AWSResourceAppSync AWSResource = "AppSync"

// AWSResourceAthena defines Athena resource
const AWSResourceAthena AWSResource = "Athena"

// AWSResourceAutoScaling defines Auto Scaling resource
const AWSResourceAutoScaling AWSResource = "Auto Scaling"

// AWSResourceAWSBatch defines AWS Batch resource
const AWSResourceAWSBatch AWSResource = "AWS Batch"

// AWSResourceBudgets defines Budgets resource
const AWSResourceBudgets AWSResource = "Budgets"

// AWSResourceCertificateManager defines Certificate Manager resource
const AWSResourceCertificateManager AWSResource = "Certificate Manager"

// AWSResourceAWSCloud9 defines AWS Cloud9 resource
const AWSResourceAWSCloud9 AWSResource = "AWS Cloud9"

// AWSResourceCloudFormation defines CloudFormation resource
const AWSResourceCloudFormation AWSResource = "CloudFormation"

// AWSResourceCloudFront defines CloudFront resource
const AWSResourceCloudFront AWSResource = "CloudFront"

// AWSResourceAWSCloudMap defines AWS Cloud Map resource
const AWSResourceAWSCloudMap AWSResource = "AWS Cloud Map"

// AWSResourceCloudTrail defines CloudTrail resource
const AWSResourceCloudTrail AWSResource = "CloudTrail"

// AWSResourceCloudWatch defines CloudWatch resource
const AWSResourceCloudWatch AWSResource = "CloudWatch"

// AWSResourceCodeBuild defines CodeBuild resource
const AWSResourceCodeBuild AWSResource = "CodeBuild"

// AWSResourceCodeCommit defines CodeCommit resource
const AWSResourceCodeCommit AWSResource = "CodeCommit"

// AWSResourceCodeDeploy defines CodeDeploy resource
const AWSResourceCodeDeploy AWSResource = "CodeDeploy"

// AWSResourceCodePipeline defines CodePipeline resource
const AWSResourceCodePipeline AWSResource = "CodePipeline"

// AWSResourceAmazonCognito defines Amazon Cognito resource
const AWSResourceAmazonCognito AWSResource = "Amazon Cognito"

// AWSResourceAWSConfig defines AWS Config resource
const AWSResourceAWSConfig AWSResource = "AWS Config"

// AWSResourceDataPipeline defines Data Pipeline resource
const AWSResourceDataPipeline AWSResource = "Data Pipeline"

// AWSResourceDAX defines DAX resource
const AWSResourceDAX AWSResource = "DAX"

// AWSResourceDirectoryService defines Directory Service resource
const AWSResourceDirectoryService AWSResource = "Directory Service"

// AWSResourceDataLifecycleManager defines Data Lifecycle Manager resource
const AWSResourceDataLifecycleManager AWSResource = "Data Lifecycle Manager"

// AWSResourceDMS defines DMS resource
const AWSResourceDMS AWSResource = "DMS"

// AWSResourceDocumentDB defines DocumentDB resource
const AWSResourceDocumentDB AWSResource = "DocumentDB"

// AWSResourceDynamoDB defines DynamoDB resource
const AWSResourceDynamoDB AWSResource = "DynamoDB"

// AWSResourceEC2 defines EC2 resource
const AWSResourceEC2 AWSResource = "EC2"

// AWSResourceECR defines ECR resource
const AWSResourceECR AWSResource = "ECR"

// AWSResourceECS defines ECS resource
const AWSResourceECS AWSResource = "ECS"

// AWSResourceEFS defines EFS resource
const AWSResourceEFS AWSResource = "EFS"

// AWSResourceEKS defines EKS resource
const AWSResourceEKS AWSResource = "EKS"

// AWSResourceElastiCache defines ElastiCache resource
const AWSResourceElastiCache AWSResource = "ElastiCache"

// AWSResourceElasticsearch defines Elasticsearch resource
const AWSResourceElasticsearch AWSResource = "Elasticsearch"

// AWSResourceElasticBeanstalk defines Elastic Beanstalk resource
const AWSResourceElasticBeanstalk AWSResource = "Elastic Beanstalk"

// AWSResourceElasticLoadBalancing defines Elastic Load Balancing resource
const AWSResourceElasticLoadBalancing AWSResource = "Elastic Load Balancing"

// AWSResourceElasticLoadBalancingV2 defines Elastic Load Balancing V2 resource
const AWSResourceElasticLoadBalancingV2 AWSResource = "Elastic Load Balancing V2"

// AWSResourceEMR defines EMR resource
const AWSResourceEMR AWSResource = "EMR"

// AWSResourceAmazonGameLift defines Amazon GameLift resource
const AWSResourceAmazonGameLift AWSResource = "Amazon GameLift"

// AWSResourceGlue defines Glue resource
const AWSResourceGlue AWSResource = "Glue"

// AWSResourceGuardDuty defines GuardDuty resource
const AWSResourceGuardDuty AWSResource = "GuardDuty"

// AWSResourceIAM defines IAM resource
const AWSResourceIAM AWSResource = "IAM"

// AWSResourceAmazonInspector defines Amazon Inspector resource
const AWSResourceAmazonInspector AWSResource = "Amazon Inspector"

// AWSResourceAWSIoT defines AWS IoT resource
const AWSResourceAWSIoT AWSResource = "AWS IoT"

// AWSResourceAWSIoT1Click defines AWS IoT 1-Click resource
const AWSResourceAWSIoT1Click AWSResource = "AWS IoT 1-Click"

// AWSResourceAWSIoTAnalytics defines AWS IoT Analytics resource
const AWSResourceAWSIoTAnalytics AWSResource = "AWS IoT Analytics"

// AWSResourceKinesis defines Kinesis resource
const AWSResourceKinesis AWSResource = "Kinesis"

// AWSResourceKMS defines KMS resource
const AWSResourceKMS AWSResource = "KMS"

// AWSResourceLambda defines Lambda resource
const AWSResourceLambda AWSResource = "Lambda"

// AWSResourceNeptune defines Neptune resource
const AWSResourceNeptune AWSResource = "Neptune"

// AWSResourceOpsWorks defines OpsWorks resource
const AWSResourceOpsWorks AWSResource = "OpsWorks"

// AWSResourceRDS defines RDS resource
const AWSResourceRDS AWSResource = "RDS"

// AWSResourceAmazonRedshift defines Amazon Redshift resource
const AWSResourceAmazonRedshift AWSResource = "Amazon Redshift"

// AWSResourceRoute53 defines Route 53 resource
const AWSResourceRoute53 AWSResource = "Route 53"

// AWSResourceS3 defines S3 resource
const AWSResourceS3 AWSResource = "S3"

// AWSResourceSageMaker defines SageMaker resource
const AWSResourceSageMaker AWSResource = "SageMaker"

// AWSResourceSecretsManager defines Secrets Manager resource
const AWSResourceSecretsManager AWSResource = "Secrets Manager"

// AWSResourceServiceCatalog defines Service Catalog resource
const AWSResourceServiceCatalog AWSResource = "Service Catalog"

// AWSResourceSES defines SES resource
const AWSResourceSES AWSResource = "SES"

// AWSResourceSimpleDB defines SimpleDB resource
const AWSResourceSimpleDB AWSResource = "SimpleDB"

// AWSResourceSNS defines SNS resource
const AWSResourceSNS AWSResource = "SNS"

// AWSResourceSQS defines SQS resource
const AWSResourceSQS AWSResource = "SQS"

// AWSResourceStepFunctions defines Step Functions resource
const AWSResourceStepFunctions AWSResource = "Step Functions"

// AWSResourceSystemsManager defines Systems Manager resource
const AWSResourceSystemsManager AWSResource = "Systems Manager"

// AWSResourceWAF defines WAF resource
const AWSResourceWAF AWSResource = "WAF"

// AWSResourceWAFRegional defines WAF Regional resource
const AWSResourceWAFRegional AWSResource = "WAF Regional"

// AWSResourceWorkSpaces defines WorkSpaces resource
const AWSResourceWorkSpaces AWSResource = "WorkSpaces"

// AWSResourceSharedPropertyTypes defines Shared Property Types resource
const AWSResourceSharedPropertyTypes AWSResource = "Shared Property Types"
