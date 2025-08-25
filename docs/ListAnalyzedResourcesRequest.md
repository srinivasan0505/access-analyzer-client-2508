

# ListAnalyzedResourcesRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**analyzerArn** | **String** | The &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources\&quot;&gt;ARN of the analyzer&lt;/a&gt; to retrieve a list of analyzed resources from. |  |
|**resourceType** | [**ResourceTypeEnum**](#ResourceTypeEnum) | The type of resource. |  [optional] |
|**nextToken** | **String** | A token used for pagination of results returned. |  [optional] |
|**maxResults** | **Integer** | The maximum number of results to return in the response. |  [optional] |



## Enum: ResourceTypeEnum

| Name | Value |
|---- | -----|
| S3_BUCKET | &quot;AWS::S3::Bucket&quot; |
| IAM_ROLE | &quot;AWS::IAM::Role&quot; |
| SQS_QUEUE | &quot;AWS::SQS::Queue&quot; |
| LAMBDA_FUNCTION | &quot;AWS::Lambda::Function&quot; |
| LAMBDA_LAYER_VERSION | &quot;AWS::Lambda::LayerVersion&quot; |
| KMS_KEY | &quot;AWS::KMS::Key&quot; |
| SECRETS_MANAGER_SECRET | &quot;AWS::SecretsManager::Secret&quot; |
| EFS_FILE_SYSTEM | &quot;AWS::EFS::FileSystem&quot; |
| EC2_SNAPSHOT | &quot;AWS::EC2::Snapshot&quot; |
| ECR_REPOSITORY | &quot;AWS::ECR::Repository&quot; |
| RDS_DB_SNAPSHOT | &quot;AWS::RDS::DBSnapshot&quot; |
| RDS_DB_CLUSTER_SNAPSHOT | &quot;AWS::RDS::DBClusterSnapshot&quot; |
| SNS_TOPIC | &quot;AWS::SNS::Topic&quot; |



