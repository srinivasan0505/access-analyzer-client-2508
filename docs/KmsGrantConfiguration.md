

# KmsGrantConfiguration

A proposed grant configuration for a KMS key. For more information, see <a href=\"https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateGrant.html\">CreateGrant</a>.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**operations** | [**List**](List.md) |  |  |
|**granteePrincipal** | [**String**](String.md) |  |  |
|**retiringPrincipal** | [**String**](String.md) |  |  [optional] |
|**constraints** | [**KmsGrantConfigurationConstraints**](KmsGrantConfigurationConstraints.md) |  |  [optional] |
|**issuingAccount** | [**String**](String.md) |  |  |



