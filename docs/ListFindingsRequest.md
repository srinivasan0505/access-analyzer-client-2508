

# ListFindingsRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**analyzerArn** | **String** | The &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources\&quot;&gt;ARN of the analyzer&lt;/a&gt; to retrieve findings from. |  |
|**filter** | [**Map&lt;String, Criterion&gt;**](Criterion.md) | A filter to match for the findings to return. |  [optional] |
|**sort** | [**ListFindingsRequestSort**](ListFindingsRequestSort.md) |  |  [optional] |
|**nextToken** | **String** | A token used for pagination of results returned. |  [optional] |
|**maxResults** | **Integer** | The maximum number of results to return in the response. |  [optional] |



