

# SecretsManagerSecretConfiguration

<p>The configuration for a Secrets Manager secret. For more information, see <a href=\"https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html\">CreateSecret</a>.</p> <p>You can propose a configuration for a new secret or an existing secret that you own by specifying the secret policy and optional KMS encryption key. If the configuration is for an existing secret and you do not specify the secret policy, the access preview uses the existing policy for the secret. If the access preview is for a new resource and you do not specify the policy, the access preview assumes a secret without a policy. To propose deletion of an existing policy, you can specify an empty string. If the proposed configuration is for a new secret and you do not specify the KMS key ID, the access preview uses the Amazon Web Services managed key <code>aws/secretsmanager</code>. If you specify an empty string for the KMS key ID, the access preview uses the Amazon Web Services managed key of the Amazon Web Services account. For more information about secret policy limits, see <a href=\"https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_limits.html\">Quotas for Secrets Manager.</a>.</p>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**kmsKeyId** | [**String**](String.md) |  |  [optional] |
|**secretPolicy** | [**String**](String.md) |  |  [optional] |



