

# CreateAccessPreviewRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**analyzerArn** | **String** | The &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources\&quot;&gt;ARN of the account analyzer&lt;/a&gt; used to generate the access preview. You can only create an access preview for analyzers with an &lt;code&gt;Account&lt;/code&gt; type and &lt;code&gt;Active&lt;/code&gt; status. |  |
|**configurations** | [**Map&lt;String, ModelConfiguration&gt;**](ModelConfiguration.md) | Access control configuration for your resource that is used to generate the access preview. The access preview includes findings for external access allowed to the resource with the proposed access control configuration. The configuration must contain exactly one element. |  |
|**clientToken** | **String** | A client token. |  [optional] |



