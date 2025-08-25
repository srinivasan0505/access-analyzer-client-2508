

# ValidatePolicyRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**locale** | [**LocaleEnum**](#LocaleEnum) | The locale to use for localizing the findings. |  [optional] |
|**policyDocument** | **String** | The JSON policy document to use as the content for the policy. |  |
|**policyType** | [**PolicyTypeEnum**](#PolicyTypeEnum) | &lt;p&gt;The type of policy to validate. Identity policies grant permissions to IAM principals. Identity policies include managed and inline policies for IAM roles, users, and groups. They also include service-control policies (SCPs) that are attached to an Amazon Web Services organization, organizational unit (OU), or an account.&lt;/p&gt; &lt;p&gt;Resource policies grant permissions on Amazon Web Services resources. Resource policies include trust policies for IAM roles and bucket policies for Amazon S3 buckets. You can provide a generic input such as identity policy or resource policy or a specific input such as managed policy or Amazon S3 bucket policy. &lt;/p&gt; |  |
|**validatePolicyResourceType** | [**ValidatePolicyResourceTypeEnum**](#ValidatePolicyResourceTypeEnum) | &lt;p&gt;The type of resource to attach to your resource policy. Specify a value for the policy validation resource type only if the policy type is &lt;code&gt;RESOURCE_POLICY&lt;/code&gt;. For example, to validate a resource policy to attach to an Amazon S3 bucket, you can choose &lt;code&gt;AWS::S3::Bucket&lt;/code&gt; for the policy validation resource type.&lt;/p&gt; &lt;p&gt;For resource types not supported as valid values, IAM Access Analyzer runs policy checks that apply to all resource policies. For example, to validate a resource policy to attach to a KMS key, do not specify a value for the policy validation resource type and IAM Access Analyzer will run policy checks that apply to all resource policies.&lt;/p&gt; |  [optional] |



## Enum: LocaleEnum

| Name | Value |
|---- | -----|
| DE | &quot;DE&quot; |
| EN | &quot;EN&quot; |
| ES | &quot;ES&quot; |
| FR | &quot;FR&quot; |
| IT | &quot;IT&quot; |
| JA | &quot;JA&quot; |
| KO | &quot;KO&quot; |
| PT_BR | &quot;PT_BR&quot; |
| ZH_CN | &quot;ZH_CN&quot; |
| ZH_TW | &quot;ZH_TW&quot; |



## Enum: PolicyTypeEnum

| Name | Value |
|---- | -----|
| IDENTITY_POLICY | &quot;IDENTITY_POLICY&quot; |
| RESOURCE_POLICY | &quot;RESOURCE_POLICY&quot; |
| SERVICE_CONTROL_POLICY | &quot;SERVICE_CONTROL_POLICY&quot; |



## Enum: ValidatePolicyResourceTypeEnum

| Name | Value |
|---- | -----|
| S3_BUCKET | &quot;AWS::S3::Bucket&quot; |
| S3_ACCESS_POINT | &quot;AWS::S3::AccessPoint&quot; |
| S3_MULTI_REGION_ACCESS_POINT | &quot;AWS::S3::MultiRegionAccessPoint&quot; |
| S3_OBJECT_LAMBDA_ACCESS_POINT | &quot;AWS::S3ObjectLambda::AccessPoint&quot; |
| IAM_ASSUME_ROLE_POLICY_DOCUMENT | &quot;AWS::IAM::AssumeRolePolicyDocument&quot; |



