

# S3AccessPointConfiguration

The configuration for an Amazon S3 access point or multi-region access point for the bucket. You can propose up to 10 access points or multi-region access points per bucket. If the proposed Amazon S3 access point configuration is for an existing bucket, the access preview uses the proposed access point configuration in place of the existing access points. To propose an access point without a policy, you can provide an empty string as the access point policy. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/creating-access-points.html\">Creating access points</a>. For more information about access point policy limits, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/access-points-restrictions-limitations.html\">Access points restrictions and limitations</a>.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**accessPointPolicy** | [**String**](String.md) |  |  [optional] |
|**publicAccessBlock** | [**S3AccessPointConfigurationPublicAccessBlock**](S3AccessPointConfigurationPublicAccessBlock.md) |  |  [optional] |
|**networkOrigin** | [**S3AccessPointConfigurationNetworkOrigin**](S3AccessPointConfigurationNetworkOrigin.md) |  |  [optional] |



