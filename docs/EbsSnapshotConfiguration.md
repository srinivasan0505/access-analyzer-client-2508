

# EbsSnapshotConfiguration

The proposed access control configuration for an Amazon EBS volume snapshot. You can propose a configuration for a new Amazon EBS volume snapshot or an Amazon EBS volume snapshot that you own by specifying the user IDs, groups, and optional KMS encryption key. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifySnapshotAttribute.html\">ModifySnapshotAttribute</a>.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**userIds** | [**List**](List.md) |  |  [optional] |
|**groups** | [**List**](List.md) |  |  [optional] |
|**kmsKeyId** | [**String**](String.md) |  |  [optional] |



