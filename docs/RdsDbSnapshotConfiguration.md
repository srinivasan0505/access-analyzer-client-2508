

# RdsDbSnapshotConfiguration

The proposed access control configuration for an Amazon RDS DB snapshot. You can propose a configuration for a new Amazon RDS DB snapshot or an Amazon RDS DB snapshot that you own by specifying the <code>RdsDbSnapshotAttributeValue</code> and optional KMS encryption key. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBSnapshotAttribute.html\">ModifyDBSnapshotAttribute</a>.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**attributes** | [**Map**](Map.md) |  |  [optional] |
|**kmsKeyId** | [**String**](String.md) |  |  [optional] |



