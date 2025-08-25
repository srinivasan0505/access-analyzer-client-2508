

# CreateAnalyzerRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**analyzerName** | **String** | The name of the analyzer to create. |  |
|**type** | [**TypeEnum**](#TypeEnum) | The type of analyzer to create. Only ACCOUNT and ORGANIZATION analyzers are supported. You can create only one analyzer per account per Region. You can create up to 5 analyzers per organization per Region. |  |
|**archiveRules** | [**List&lt;InlineArchiveRule&gt;**](InlineArchiveRule.md) | Specifies the archive rules to add for the analyzer. Archive rules automatically archive findings that meet the criteria you define for the rule. |  [optional] |
|**tags** | **Map&lt;String, String&gt;** | The tags to apply to the analyzer. |  [optional] |
|**clientToken** | **String** | A client token. |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| ACCOUNT | &quot;ACCOUNT&quot; |
| ORGANIZATION | &quot;ORGANIZATION&quot; |



