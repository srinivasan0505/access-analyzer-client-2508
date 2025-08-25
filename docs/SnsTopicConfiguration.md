

# SnsTopicConfiguration

The proposed access control configuration for an Amazon SNS topic. You can propose a configuration for a new Amazon SNS topic or an existing Amazon SNS topic that you own by specifying the policy. If the configuration is for an existing Amazon SNS topic and you do not specify the Amazon SNS policy, then the access preview uses the existing Amazon SNS policy for the topic. If the access preview is for a new resource and you do not specify the policy, then the access preview assumes an Amazon SNS topic without a policy. To propose deletion of an existing Amazon SNS topic policy, you can specify an empty string for the Amazon SNS policy. For more information, see <a href=\"https://docs.aws.amazon.com/sns/latest/api/API_Topic.html\">Topic</a>.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**topicPolicy** | [**String**](String.md) |  |  [optional] |



