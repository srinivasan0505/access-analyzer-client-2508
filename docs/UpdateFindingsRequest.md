

# UpdateFindingsRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**analyzerArn** | **String** | The &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources\&quot;&gt;ARN of the analyzer&lt;/a&gt; that generated the findings to update. |  |
|**status** | [**StatusEnum**](#StatusEnum) | The state represents the action to take to update the finding Status. Use &lt;code&gt;ARCHIVE&lt;/code&gt; to change an Active finding to an Archived finding. Use &lt;code&gt;ACTIVE&lt;/code&gt; to change an Archived finding to an Active finding. |  |
|**ids** | **List&lt;String&gt;** | The IDs of the findings to update. |  [optional] |
|**resourceArn** | **String** | The ARN of the resource identified in the finding. |  [optional] |
|**clientToken** | **String** | A client token. |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| ACTIVE | &quot;ACTIVE&quot; |
| ARCHIVED | &quot;ARCHIVED&quot; |



