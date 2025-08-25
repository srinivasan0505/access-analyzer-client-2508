package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedPolicyResult represents the GeneratedPolicyResult schema from the OpenAPI specification
type GeneratedPolicyResult struct {
	Generatedpolicies interface{} `json:"generatedPolicies,omitempty"`
	Properties interface{} `json:"properties"` // A <code>GeneratedPolicyProperties</code> object that contains properties of the generated policy.
}

// S3BucketAclGrantConfiguration represents the S3BucketAclGrantConfiguration schema from the OpenAPI specification
type S3BucketAclGrantConfiguration struct {
	Grantee S3BucketAclGrantConfigurationgrantee `json:"grantee"`
	Permission interface{} `json:"permission"`
}

// InternetConfiguration represents the InternetConfiguration schema from the OpenAPI specification
type InternetConfiguration struct {
}

// ConfigurationsqsQueue represents the ConfigurationsqsQueue schema from the OpenAPI specification
type ConfigurationsqsQueue struct {
	Queuepolicy interface{} `json:"queuePolicy,omitempty"`
}

// CreateAccessPreviewResponse represents the CreateAccessPreviewResponse schema from the OpenAPI specification
type CreateAccessPreviewResponse struct {
	Id interface{} `json:"id"`
}

// ListFindingsRequest represents the ListFindingsRequest schema from the OpenAPI specification
type ListFindingsRequest struct {
	Analyzerarn interface{} `json:"analyzerArn"`
	Filter interface{} `json:"filter,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Sort ListFindingsRequestsort `json:"sort,omitempty"`
}

// GetGeneratedPolicyResponsejobDetails represents the GetGeneratedPolicyResponsejobDetails schema from the OpenAPI specification
type GetGeneratedPolicyResponsejobDetails struct {
	Completedon interface{} `json:"completedOn,omitempty"`
	Joberror JobDetailsjobError `json:"jobError,omitempty"`
	Jobid interface{} `json:"jobId"`
	Startedon interface{} `json:"startedOn"`
	Status interface{} `json:"status"`
}

// ListAccessPreviewFindingsrequest represents the ListAccessPreviewFindingsrequest schema from the OpenAPI specification
type ListAccessPreviewFindingsrequest struct {
	Filter map[string]interface{} `json:"filter,omitempty"` // Criteria to filter the returned findings.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results to return in the response.
	Nexttoken string `json:"nextToken,omitempty"` // A token used for pagination of results returned.
	Analyzerarn string `json:"analyzerArn"` // The <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources">ARN of the analyzer</a> used to generate the access.
}

// S3BucketAclGrantConfigurationgrantee represents the S3BucketAclGrantConfigurationgrantee schema from the OpenAPI specification
type S3BucketAclGrantConfigurationgrantee struct {
	Id interface{} `json:"id,omitempty"`
	Uri interface{} `json:"uri,omitempty"`
}

// Span represents the Span schema from the OpenAPI specification
type Span struct {
	End Spanend `json:"end"`
	Start Spanstart `json:"start"`
}

// ListFindingsResponse represents the ListFindingsResponse schema from the OpenAPI specification
type ListFindingsResponse struct {
	Findings interface{} `json:"findings"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListAnalyzedResourcesrequest represents the ListAnalyzedResourcesrequest schema from the OpenAPI specification
type ListAnalyzedResourcesrequest struct {
	Analyzerarn string `json:"analyzerArn"` // The <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources">ARN of the analyzer</a> to retrieve a list of analyzed resources from.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results to return in the response.
	Nexttoken string `json:"nextToken,omitempty"` // A token used for pagination of results returned.
	Resourcetype string `json:"resourceType,omitempty"` // The type of resource.
}

// JobDetailsjobError represents the JobDetailsjobError schema from the OpenAPI specification
type JobDetailsjobError struct {
	Code interface{} `json:"code"`
	Message interface{} `json:"message"`
}

// ValidatePolicyrequest represents the ValidatePolicyrequest schema from the OpenAPI specification
type ValidatePolicyrequest struct {
	Locale string `json:"locale,omitempty"` // The locale to use for localizing the findings.
	Policydocument string `json:"policyDocument"` // The JSON policy document to use as the content for the policy.
	Policytype string `json:"policyType"` // <p>The type of policy to validate. Identity policies grant permissions to IAM principals. Identity policies include managed and inline policies for IAM roles, users, and groups. They also include service-control policies (SCPs) that are attached to an Amazon Web Services organization, organizational unit (OU), or an account.</p> <p>Resource policies grant permissions on Amazon Web Services resources. Resource policies include trust policies for IAM roles and bucket policies for Amazon S3 buckets. You can provide a generic input such as identity policy or resource policy or a specific input such as managed policy or Amazon S3 bucket policy. </p>
	Validatepolicyresourcetype string `json:"validatePolicyResourceType,omitempty"` // <p>The type of resource to attach to your resource policy. Specify a value for the policy validation resource type only if the policy type is <code>RESOURCE_POLICY</code>. For example, to validate a resource policy to attach to an Amazon S3 bucket, you can choose <code>AWS::S3::Bucket</code> for the policy validation resource type.</p> <p>For resource types not supported as valid values, IAM Access Analyzer runs policy checks that apply to all resource policies. For example, to validate a resource policy to attach to a KMS key, do not specify a value for the policy validation resource type and IAM Access Analyzer will run policy checks that apply to all resource policies.</p>
}

// ListAnalyzedResourcesRequest represents the ListAnalyzedResourcesRequest schema from the OpenAPI specification
type ListAnalyzedResourcesRequest struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Resourcetype interface{} `json:"resourceType,omitempty"`
	Analyzerarn interface{} `json:"analyzerArn"`
	Maxresults interface{} `json:"maxResults,omitempty"`
}

// PrincipalMap represents the PrincipalMap schema from the OpenAPI specification
type PrincipalMap struct {
}

// ValidatePolicyResponse represents the ValidatePolicyResponse schema from the OpenAPI specification
type ValidatePolicyResponse struct {
	Findings interface{} `json:"findings"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// ArchiveRuleSummary represents the ArchiveRuleSummary schema from the OpenAPI specification
type ArchiveRuleSummary struct {
	Createdat interface{} `json:"createdAt"`
	Filter interface{} `json:"filter"`
	Rulename interface{} `json:"ruleName"`
	Updatedat interface{} `json:"updatedAt"`
}

// Locationspan represents the Locationspan schema from the OpenAPI specification
type Locationspan struct {
	End Spanend `json:"end"`
	Start Spanstart `json:"start"`
}

// CancelPolicyGenerationRequest represents the CancelPolicyGenerationRequest schema from the OpenAPI specification
type CancelPolicyGenerationRequest struct {
}

// RdsDbClusterSnapshotAttributesMap represents the RdsDbClusterSnapshotAttributesMap schema from the OpenAPI specification
type RdsDbClusterSnapshotAttributesMap struct {
}

// ListFindingsrequestsort represents the ListFindingsrequestsort schema from the OpenAPI specification
type ListFindingsrequestsort struct {
	Orderby interface{} `json:"orderBy,omitempty"`
	Attributename interface{} `json:"attributeName,omitempty"`
}

// GetAccessPreviewRequest represents the GetAccessPreviewRequest schema from the OpenAPI specification
type GetAccessPreviewRequest struct {
}

// TagsMap represents the TagsMap schema from the OpenAPI specification
type TagsMap struct {
}

// CreateArchiveRuleRequest represents the CreateArchiveRuleRequest schema from the OpenAPI specification
type CreateArchiveRuleRequest struct {
	Rulename interface{} `json:"ruleName"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Filter interface{} `json:"filter"`
}

// TagResourcerequest represents the TagResourcerequest schema from the OpenAPI specification
type TagResourcerequest struct {
	Tags map[string]interface{} `json:"tags"` // The tags to add to the resource.
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// S3AccessPointConfigurationsMap represents the S3AccessPointConfigurationsMap schema from the OpenAPI specification
type S3AccessPointConfigurationsMap struct {
}

// KmsGrantConfiguration represents the KmsGrantConfiguration schema from the OpenAPI specification
type KmsGrantConfiguration struct {
	Retiringprincipal interface{} `json:"retiringPrincipal,omitempty"`
	Constraints KmsGrantConfigurationconstraints `json:"constraints,omitempty"`
	Granteeprincipal interface{} `json:"granteePrincipal"`
	Issuingaccount interface{} `json:"issuingAccount"`
	Operations interface{} `json:"operations"`
}

// ListArchiveRulesRequest represents the ListArchiveRulesRequest schema from the OpenAPI specification
type ListArchiveRulesRequest struct {
}

// FindingSourceDetail represents the FindingSourceDetail schema from the OpenAPI specification
type FindingSourceDetail struct {
	Accesspointaccount interface{} `json:"accessPointAccount,omitempty"`
	Accesspointarn interface{} `json:"accessPointArn,omitempty"`
}

// GetAnalyzerResponseanalyzer represents the GetAnalyzerResponseanalyzer schema from the OpenAPI specification
type GetAnalyzerResponseanalyzer struct {
	Name interface{} `json:"name"`
	Status interface{} `json:"status"`
	Arn interface{} `json:"arn"`
	Createdat interface{} `json:"createdAt"`
	Lastresourceanalyzed interface{} `json:"lastResourceAnalyzed,omitempty"`
	Lastresourceanalyzedat interface{} `json:"lastResourceAnalyzedAt,omitempty"`
	Statusreason AnalyzerSummarystatusReason `json:"statusReason,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	TypeField interface{} `json:"type"`
}

// RdsDbClusterSnapshotAttributeValue represents the RdsDbClusterSnapshotAttributeValue schema from the OpenAPI specification
type RdsDbClusterSnapshotAttributeValue struct {
	Accountids interface{} `json:"accountIds,omitempty"`
}

// StartPolicyGenerationResponse represents the StartPolicyGenerationResponse schema from the OpenAPI specification
type StartPolicyGenerationResponse struct {
	Jobid interface{} `json:"jobId"`
}

// StartResourceScanrequest represents the StartResourceScanrequest schema from the OpenAPI specification
type StartResourceScanrequest struct {
	Resourcearn string `json:"resourceArn"` // The ARN of the resource to scan.
	Resourceowneraccount string `json:"resourceOwnerAccount,omitempty"` // The Amazon Web Services account ID that owns the resource. For most Amazon Web Services resources, the owning account is the account in which the resource was created.
	Analyzerarn string `json:"analyzerArn"` // The <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources">ARN of the analyzer</a> to use to scan the policies applied to the specified resource.
}

// EbsSnapshotConfiguration represents the EbsSnapshotConfiguration schema from the OpenAPI specification
type EbsSnapshotConfiguration struct {
	Groups interface{} `json:"groups,omitempty"`
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
	Userids interface{} `json:"userIds,omitempty"`
}

// ConfigurationiamRole represents the ConfigurationiamRole schema from the OpenAPI specification
type ConfigurationiamRole struct {
	Trustpolicy interface{} `json:"trustPolicy,omitempty"`
}

// EfsFileSystemConfiguration represents the EfsFileSystemConfiguration schema from the OpenAPI specification
type EfsFileSystemConfiguration struct {
	Filesystempolicy interface{} `json:"fileSystemPolicy,omitempty"`
}

// CloudTrailProperties represents the CloudTrailProperties schema from the OpenAPI specification
type CloudTrailProperties struct {
	Trailproperties interface{} `json:"trailProperties"`
	Endtime interface{} `json:"endTime"`
	Starttime interface{} `json:"startTime"`
}

// GetGeneratedPolicyResponsegeneratedPolicyResult represents the GetGeneratedPolicyResponsegeneratedPolicyResult schema from the OpenAPI specification
type GetGeneratedPolicyResponsegeneratedPolicyResult struct {
	Generatedpolicies interface{} `json:"generatedPolicies,omitempty"`
	Properties interface{} `json:"properties"` // A <code>GeneratedPolicyProperties</code> object that contains properties of the generated policy.
}

// SqsQueueConfiguration represents the SqsQueueConfiguration schema from the OpenAPI specification
type SqsQueueConfiguration struct {
	Queuepolicy interface{} `json:"queuePolicy,omitempty"`
}

// StartResourceScanRequest represents the StartResourceScanRequest schema from the OpenAPI specification
type StartResourceScanRequest struct {
	Analyzerarn interface{} `json:"analyzerArn"`
	Resourcearn interface{} `json:"resourceArn"`
	Resourceowneraccount interface{} `json:"resourceOwnerAccount,omitempty"`
}

// ValidatePolicyFinding represents the ValidatePolicyFinding schema from the OpenAPI specification
type ValidatePolicyFinding struct {
	Findingdetails interface{} `json:"findingDetails"`
	Findingtype interface{} `json:"findingType"`
	Issuecode interface{} `json:"issueCode"`
	Learnmorelink interface{} `json:"learnMoreLink"`
	Locations interface{} `json:"locations"`
}

// GetAnalyzedResourceResponse represents the GetAnalyzedResourceResponse schema from the OpenAPI specification
type GetAnalyzedResourceResponse struct {
	Resource GetAnalyzedResourceResponseresource `json:"resource,omitempty"`
}

// AccessPreviewFinding represents the AccessPreviewFinding schema from the OpenAPI specification
type AccessPreviewFinding struct {
	Existingfindingid interface{} `json:"existingFindingId,omitempty"`
	Action interface{} `json:"action,omitempty"`
	Createdat interface{} `json:"createdAt"`
	Condition interface{} `json:"condition,omitempty"`
	Resourceowneraccount interface{} `json:"resourceOwnerAccount"`
	Resourcetype interface{} `json:"resourceType"`
	Status interface{} `json:"status"`
	Resource interface{} `json:"resource,omitempty"`
	Id interface{} `json:"id"`
	Principal interface{} `json:"principal,omitempty"`
	Changetype interface{} `json:"changeType"`
	Ispublic interface{} `json:"isPublic,omitempty"`
	ErrorField interface{} `json:"error,omitempty"`
	Existingfindingstatus interface{} `json:"existingFindingStatus,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
}

// Configurations3Bucket represents the Configurations3Bucket schema from the OpenAPI specification
type Configurations3Bucket struct {
	Bucketpolicy interface{} `json:"bucketPolicy,omitempty"`
	Bucketpublicaccessblock S3BucketConfigurationbucketPublicAccessBlock `json:"bucketPublicAccessBlock,omitempty"`
	Accesspoints interface{} `json:"accessPoints,omitempty"`
	Bucketaclgrants interface{} `json:"bucketAclGrants,omitempty"`
}

// ConfigurationefsFileSystem represents the ConfigurationefsFileSystem schema from the OpenAPI specification
type ConfigurationefsFileSystem struct {
	Filesystempolicy interface{} `json:"fileSystemPolicy,omitempty"`
}

// SortCriteria represents the SortCriteria schema from the OpenAPI specification
type SortCriteria struct {
	Attributename interface{} `json:"attributeName,omitempty"`
	Orderby interface{} `json:"orderBy,omitempty"`
}

// EcrRepositoryConfiguration represents the EcrRepositoryConfiguration schema from the OpenAPI specification
type EcrRepositoryConfiguration struct {
	Repositorypolicy interface{} `json:"repositoryPolicy,omitempty"`
}

// ApplyArchiveRuleRequest represents the ApplyArchiveRuleRequest schema from the OpenAPI specification
type ApplyArchiveRuleRequest struct {
	Analyzerarn interface{} `json:"analyzerArn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Rulename interface{} `json:"ruleName"`
}

// StartPolicyGenerationRequestcloudTrailDetails represents the StartPolicyGenerationRequestcloudTrailDetails schema from the OpenAPI specification
type StartPolicyGenerationRequestcloudTrailDetails struct {
	Accessrole interface{} `json:"accessRole"`
	Endtime interface{} `json:"endTime,omitempty"`
	Starttime interface{} `json:"startTime"`
	Trails interface{} `json:"trails"`
}

// Substring represents the Substring schema from the OpenAPI specification
type Substring struct {
	Length interface{} `json:"length"`
	Start interface{} `json:"start"`
}

// ApplyArchiveRulerequest represents the ApplyArchiveRulerequest schema from the OpenAPI specification
type ApplyArchiveRulerequest struct {
	Rulename string `json:"ruleName"` // The name of the rule to apply.
	Analyzerarn string `json:"analyzerArn"` // The Amazon resource name (ARN) of the analyzer.
	Clienttoken string `json:"clientToken,omitempty"` // A client token.
}

// AccessPreviewStatusReason represents the AccessPreviewStatusReason schema from the OpenAPI specification
type AccessPreviewStatusReason struct {
	Code interface{} `json:"code"`
}

// GeneratedPolicyProperties represents the GeneratedPolicyProperties schema from the OpenAPI specification
type GeneratedPolicyProperties struct {
	Iscomplete interface{} `json:"isComplete,omitempty"`
	Principalarn interface{} `json:"principalArn"`
	Cloudtrailproperties GeneratedPolicyPropertiescloudTrailProperties `json:"cloudTrailProperties,omitempty"`
}

// IamRoleConfiguration represents the IamRoleConfiguration schema from the OpenAPI specification
type IamRoleConfiguration struct {
	Trustpolicy interface{} `json:"trustPolicy,omitempty"`
}

// GetGeneratedPolicyRequest represents the GetGeneratedPolicyRequest schema from the OpenAPI specification
type GetGeneratedPolicyRequest struct {
}

// RdsDbClusterSnapshotConfiguration represents the RdsDbClusterSnapshotConfiguration schema from the OpenAPI specification
type RdsDbClusterSnapshotConfiguration struct {
	Attributes interface{} `json:"attributes,omitempty"`
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
}

// GetArchiveRuleResponse represents the GetArchiveRuleResponse schema from the OpenAPI specification
type GetArchiveRuleResponse struct {
	Archiverule ArchiveRuleSummary `json:"archiveRule"` // Contains information about an archive rule.
}

// GeneratedPolicy represents the GeneratedPolicy schema from the OpenAPI specification
type GeneratedPolicy struct {
	Policy interface{} `json:"policy"`
}

// CreateAccessPreviewRequest represents the CreateAccessPreviewRequest schema from the OpenAPI specification
type CreateAccessPreviewRequest struct {
	Configurations interface{} `json:"configurations"`
	Analyzerarn interface{} `json:"analyzerArn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// KmsKeyPoliciesMap represents the KmsKeyPoliciesMap schema from the OpenAPI specification
type KmsKeyPoliciesMap struct {
}

// ConfigurationsnsTopic represents the ConfigurationsnsTopic schema from the OpenAPI specification
type ConfigurationsnsTopic struct {
	Topicpolicy interface{} `json:"topicPolicy,omitempty"`
}

// RdsDbSnapshotConfiguration represents the RdsDbSnapshotConfiguration schema from the OpenAPI specification
type RdsDbSnapshotConfiguration struct {
	Attributes interface{} `json:"attributes,omitempty"`
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
}

// ConfigurationebsSnapshot represents the ConfigurationebsSnapshot schema from the OpenAPI specification
type ConfigurationebsSnapshot struct {
	Groups interface{} `json:"groups,omitempty"`
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
	Userids interface{} `json:"userIds,omitempty"`
}

// GetAccessPreviewResponse represents the GetAccessPreviewResponse schema from the OpenAPI specification
type GetAccessPreviewResponse struct {
	Accesspreview GetAccessPreviewResponseaccessPreview `json:"accessPreview"`
}

// RdsDbSnapshotAttributeValue represents the RdsDbSnapshotAttributeValue schema from the OpenAPI specification
type RdsDbSnapshotAttributeValue struct {
	Accountids interface{} `json:"accountIds,omitempty"`
}

// GetFindingResponsefinding represents the GetFindingResponsefinding schema from the OpenAPI specification
type GetFindingResponsefinding struct {
	Condition interface{} `json:"condition"`
	Resourceowneraccount interface{} `json:"resourceOwnerAccount"`
	Sources interface{} `json:"sources,omitempty"`
	Analyzedat interface{} `json:"analyzedAt"`
	Createdat interface{} `json:"createdAt"`
	ErrorField interface{} `json:"error,omitempty"`
	Resourcetype interface{} `json:"resourceType"`
	Status interface{} `json:"status"`
	Ispublic interface{} `json:"isPublic,omitempty"`
	Principal interface{} `json:"principal,omitempty"`
	Updatedat interface{} `json:"updatedAt"`
	Action interface{} `json:"action,omitempty"`
	Id interface{} `json:"id"`
	Resource interface{} `json:"resource,omitempty"`
}

// AnalyzedResourceSummary represents the AnalyzedResourceSummary schema from the OpenAPI specification
type AnalyzedResourceSummary struct {
	Resourceowneraccount interface{} `json:"resourceOwnerAccount"`
	Resourcetype interface{} `json:"resourceType"`
	Resourcearn interface{} `json:"resourceArn"`
}

// CreateAnalyzerRequest represents the CreateAnalyzerRequest schema from the OpenAPI specification
type CreateAnalyzerRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	TypeField interface{} `json:"type"`
	Analyzername interface{} `json:"analyzerName"`
	Archiverules interface{} `json:"archiveRules,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// Criterion represents the Criterion schema from the OpenAPI specification
type Criterion struct {
	Neq interface{} `json:"neq,omitempty"`
	Contains interface{} `json:"contains,omitempty"`
	Eq interface{} `json:"eq,omitempty"`
	Exists interface{} `json:"exists,omitempty"`
}

// JobDetails represents the JobDetails schema from the OpenAPI specification
type JobDetails struct {
	Status interface{} `json:"status"`
	Completedon interface{} `json:"completedOn,omitempty"`
	Joberror JobDetailsjobError `json:"jobError,omitempty"`
	Jobid interface{} `json:"jobId"`
	Startedon interface{} `json:"startedOn"`
}

// StartPolicyGenerationrequestcloudTrailDetails represents the StartPolicyGenerationrequestcloudTrailDetails schema from the OpenAPI specification
type StartPolicyGenerationrequestcloudTrailDetails struct {
	Starttime interface{} `json:"startTime,omitempty"`
	Trails interface{} `json:"trails,omitempty"`
	Accessrole interface{} `json:"accessRole,omitempty"`
	Endtime interface{} `json:"endTime,omitempty"`
}

// VpcConfiguration represents the VpcConfiguration schema from the OpenAPI specification
type VpcConfiguration struct {
	Vpcid interface{} `json:"vpcId"`
}

// Position represents the Position schema from the OpenAPI specification
type Position struct {
	Column interface{} `json:"column"`
	Line interface{} `json:"line"`
	Offset interface{} `json:"offset"`
}

// ListPolicyGenerationsRequest represents the ListPolicyGenerationsRequest schema from the OpenAPI specification
type ListPolicyGenerationsRequest struct {
}

// GetFindingResponse represents the GetFindingResponse schema from the OpenAPI specification
type GetFindingResponse struct {
	Finding GetFindingResponsefinding `json:"finding,omitempty"`
}

// CloudTrailDetails represents the CloudTrailDetails schema from the OpenAPI specification
type CloudTrailDetails struct {
	Starttime interface{} `json:"startTime"`
	Trails interface{} `json:"trails"`
	Accessrole interface{} `json:"accessRole"`
	Endtime interface{} `json:"endTime,omitempty"`
}

// NetworkOriginConfiguration represents the NetworkOriginConfiguration schema from the OpenAPI specification
type NetworkOriginConfiguration struct {
	Internetconfiguration interface{} `json:"internetConfiguration,omitempty"`
	Vpcconfiguration VpcConfiguration `json:"vpcConfiguration,omitempty"` // The proposed virtual private cloud (VPC) configuration for the Amazon S3 access point. VPC configuration does not apply to multi-region access points. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_VpcConfiguration.html">VpcConfiguration</a>.
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// Spanend represents the Spanend schema from the OpenAPI specification
type Spanend struct {
	Offset interface{} `json:"offset"`
	Column interface{} `json:"column"`
	Line interface{} `json:"line"`
}

// ConditionKeyMap represents the ConditionKeyMap schema from the OpenAPI specification
type ConditionKeyMap struct {
}

// UpdateFindingsRequest represents the UpdateFindingsRequest schema from the OpenAPI specification
type UpdateFindingsRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Ids interface{} `json:"ids,omitempty"`
	Resourcearn interface{} `json:"resourceArn,omitempty"`
	Status interface{} `json:"status"`
	Analyzerarn interface{} `json:"analyzerArn"`
}

// FilterCriteriaMap represents the FilterCriteriaMap schema from the OpenAPI specification
type FilterCriteriaMap struct {
}

// StartPolicyGenerationrequest represents the StartPolicyGenerationrequest schema from the OpenAPI specification
type StartPolicyGenerationrequest struct {
	Clienttoken string `json:"clientToken,omitempty"` // <p>A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. Idempotency ensures that an API request completes only once. With an idempotent request, if the original request completes successfully, the subsequent retries with the same client token return the result from the original successful request and they have no additional effect.</p> <p>If you do not specify a client token, one is automatically generated by the Amazon Web Services SDK.</p>
	Cloudtraildetails StartPolicyGenerationrequestcloudTrailDetails `json:"cloudTrailDetails,omitempty"` // Contains information about CloudTrail access.
	Policygenerationdetails StartPolicyGenerationrequestpolicyGenerationDetails `json:"policyGenerationDetails"` // Contains the ARN details about the IAM entity for which the policy is generated.
}

// StartPolicyGenerationRequest represents the StartPolicyGenerationRequest schema from the OpenAPI specification
type StartPolicyGenerationRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Cloudtraildetails StartPolicyGenerationRequestcloudTrailDetails `json:"cloudTrailDetails,omitempty"`
	Policygenerationdetails StartPolicyGenerationRequestpolicyGenerationDetails `json:"policyGenerationDetails"`
}

// ListAccessPreviewFindingsResponse represents the ListAccessPreviewFindingsResponse schema from the OpenAPI specification
type ListAccessPreviewFindingsResponse struct {
	Findings interface{} `json:"findings"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateArchiveRuleRequest represents the UpdateArchiveRuleRequest schema from the OpenAPI specification
type UpdateArchiveRuleRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Filter interface{} `json:"filter"`
}

// CreateArchiveRulerequest represents the CreateArchiveRulerequest schema from the OpenAPI specification
type CreateArchiveRulerequest struct {
	Filter map[string]interface{} `json:"filter"` // The criteria for the rule.
	Rulename string `json:"ruleName"` // The name of the rule to create.
	Clienttoken string `json:"clientToken,omitempty"` // A client token.
}

// StartPolicyGenerationrequestpolicyGenerationDetails represents the StartPolicyGenerationrequestpolicyGenerationDetails schema from the OpenAPI specification
type StartPolicyGenerationrequestpolicyGenerationDetails struct {
	Principalarn interface{} `json:"principalArn,omitempty"`
}

// StartPolicyGenerationRequestpolicyGenerationDetails represents the StartPolicyGenerationRequestpolicyGenerationDetails schema from the OpenAPI specification
type StartPolicyGenerationRequestpolicyGenerationDetails struct {
	Principalarn interface{} `json:"principalArn"`
}

// ValidatePolicyRequest represents the ValidatePolicyRequest schema from the OpenAPI specification
type ValidatePolicyRequest struct {
	Locale interface{} `json:"locale,omitempty"`
	Policydocument interface{} `json:"policyDocument"`
	Policytype interface{} `json:"policyType"`
	Validatepolicyresourcetype interface{} `json:"validatePolicyResourceType,omitempty"`
}

// CreateAnalyzerResponse represents the CreateAnalyzerResponse schema from the OpenAPI specification
type CreateAnalyzerResponse struct {
	Arn interface{} `json:"arn,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// KmsKeyConfiguration represents the KmsKeyConfiguration schema from the OpenAPI specification
type KmsKeyConfiguration struct {
	Keypolicies interface{} `json:"keyPolicies,omitempty"`
	Grants interface{} `json:"grants,omitempty"`
}

// GetAccessPreviewResponseaccessPreview represents the GetAccessPreviewResponseaccessPreview schema from the OpenAPI specification
type GetAccessPreviewResponseaccessPreview struct {
	Analyzerarn interface{} `json:"analyzerArn"`
	Configurations interface{} `json:"configurations"`
	Createdat interface{} `json:"createdAt"`
	Id interface{} `json:"id"`
	Status interface{} `json:"status"`
	Statusreason AccessPreviewstatusReason `json:"statusReason,omitempty"`
}

// AccessPreviewstatusReason represents the AccessPreviewstatusReason schema from the OpenAPI specification
type AccessPreviewstatusReason struct {
	Code interface{} `json:"code"`
}

// ConfigurationsecretsManagerSecret represents the ConfigurationsecretsManagerSecret schema from the OpenAPI specification
type ConfigurationsecretsManagerSecret struct {
	Secretpolicy interface{} `json:"secretPolicy,omitempty"`
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
}

// CreateAnalyzerrequest represents the CreateAnalyzerrequest schema from the OpenAPI specification
type CreateAnalyzerrequest struct {
	Clienttoken string `json:"clientToken,omitempty"` // A client token.
	Tags map[string]interface{} `json:"tags,omitempty"` // The tags to apply to the analyzer.
	TypeField string `json:"type"` // The type of analyzer to create. Only ACCOUNT and ORGANIZATION analyzers are supported. You can create only one analyzer per account per Region. You can create up to 5 analyzers per organization per Region.
	Analyzername string `json:"analyzerName"` // The name of the analyzer to create.
	Archiverules []InlineArchiveRule `json:"archiveRules,omitempty"` // Specifies the archive rules to add for the analyzer. Archive rules automatically archive findings that meet the criteria you define for the rule.
}

// ConfigurationsMap represents the ConfigurationsMap schema from the OpenAPI specification
type ConfigurationsMap struct {
}

// UpdateArchiveRulerequest represents the UpdateArchiveRulerequest schema from the OpenAPI specification
type UpdateArchiveRulerequest struct {
	Filter map[string]interface{} `json:"filter"` // A filter to match for the rules to update. Only rules that match the filter are updated.
	Clienttoken string `json:"clientToken,omitempty"` // A client token.
}

// GetArchiveRuleRequest represents the GetArchiveRuleRequest schema from the OpenAPI specification
type GetArchiveRuleRequest struct {
}

// KmsConstraintsMap represents the KmsConstraintsMap schema from the OpenAPI specification
type KmsConstraintsMap struct {
}

// ListArchiveRulesResponse represents the ListArchiveRulesResponse schema from the OpenAPI specification
type ListArchiveRulesResponse struct {
	Archiverules interface{} `json:"archiveRules"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// FindingSourcedetail represents the FindingSourcedetail schema from the OpenAPI specification
type FindingSourcedetail struct {
	Accesspointaccount interface{} `json:"accessPointAccount,omitempty"`
	Accesspointarn interface{} `json:"accessPointArn,omitempty"`
}

// StatusReason represents the StatusReason schema from the OpenAPI specification
type StatusReason struct {
	Code interface{} `json:"code"`
}

// CancelPolicyGenerationResponse represents the CancelPolicyGenerationResponse schema from the OpenAPI specification
type CancelPolicyGenerationResponse struct {
}

// S3AccessPointConfiguration represents the S3AccessPointConfiguration schema from the OpenAPI specification
type S3AccessPointConfiguration struct {
	Accesspointpolicy interface{} `json:"accessPointPolicy,omitempty"`
	Networkorigin S3AccessPointConfigurationnetworkOrigin `json:"networkOrigin,omitempty"`
	Publicaccessblock S3AccessPointConfigurationpublicAccessBlock `json:"publicAccessBlock,omitempty"`
}

// AnalyzerSummarystatusReason represents the AnalyzerSummarystatusReason schema from the OpenAPI specification
type AnalyzerSummarystatusReason struct {
	Code interface{} `json:"code"`
}

// UpdateFindingsrequest represents the UpdateFindingsrequest schema from the OpenAPI specification
type UpdateFindingsrequest struct {
	Analyzerarn string `json:"analyzerArn"` // The <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources">ARN of the analyzer</a> that generated the findings to update.
	Clienttoken string `json:"clientToken,omitempty"` // A client token.
	Ids []string `json:"ids,omitempty"` // The IDs of the findings to update.
	Resourcearn string `json:"resourceArn,omitempty"` // The ARN of the resource identified in the finding.
	Status string `json:"status"` // The state represents the action to take to update the finding Status. Use <code>ARCHIVE</code> to change an Active finding to an Archived finding. Use <code>ACTIVE</code> to change an Archived finding to an Active finding.
}

// DeleteAnalyzerRequest represents the DeleteAnalyzerRequest schema from the OpenAPI specification
type DeleteAnalyzerRequest struct {
}

// SnsTopicConfiguration represents the SnsTopicConfiguration schema from the OpenAPI specification
type SnsTopicConfiguration struct {
	Topicpolicy interface{} `json:"topicPolicy,omitempty"`
}

// AccessPreviewSummary represents the AccessPreviewSummary schema from the OpenAPI specification
type AccessPreviewSummary struct {
	Analyzerarn interface{} `json:"analyzerArn"`
	Createdat interface{} `json:"createdAt"`
	Id interface{} `json:"id"`
	Status interface{} `json:"status"`
	Statusreason AccessPreviewStatusReason `json:"statusReason,omitempty"` // Provides more details about the current status of the access preview. For example, if the creation of the access preview fails, a <code>Failed</code> status is returned. This failure can be due to an internal issue with the analysis or due to an invalid proposed resource configuration.
}

// AnalyzerSummary represents the AnalyzerSummary schema from the OpenAPI specification
type AnalyzerSummary struct {
	Name interface{} `json:"name"`
	Status interface{} `json:"status"`
	Arn interface{} `json:"arn"`
	Createdat interface{} `json:"createdAt"`
	Lastresourceanalyzed interface{} `json:"lastResourceAnalyzed,omitempty"`
	Lastresourceanalyzedat interface{} `json:"lastResourceAnalyzedAt,omitempty"`
	Statusreason AnalyzerSummarystatusReason `json:"statusReason,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	TypeField interface{} `json:"type"`
}

// TrailProperties represents the TrailProperties schema from the OpenAPI specification
type TrailProperties struct {
	Regions interface{} `json:"regions,omitempty"`
	Allregions interface{} `json:"allRegions,omitempty"`
	Cloudtrailarn interface{} `json:"cloudTrailArn"`
}

// ConfigurationecrRepository represents the ConfigurationecrRepository schema from the OpenAPI specification
type ConfigurationecrRepository struct {
	Repositorypolicy interface{} `json:"repositoryPolicy,omitempty"`
}

// S3AccessPointConfigurationpublicAccessBlock represents the S3AccessPointConfigurationpublicAccessBlock schema from the OpenAPI specification
type S3AccessPointConfigurationpublicAccessBlock struct {
	Ignorepublicacls interface{} `json:"ignorePublicAcls"`
	Restrictpublicbuckets interface{} `json:"restrictPublicBuckets"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// AccessPreview represents the AccessPreview schema from the OpenAPI specification
type AccessPreview struct {
	Id interface{} `json:"id"`
	Status interface{} `json:"status"`
	Statusreason AccessPreviewstatusReason `json:"statusReason,omitempty"`
	Analyzerarn interface{} `json:"analyzerArn"`
	Configurations interface{} `json:"configurations"`
	Createdat interface{} `json:"createdAt"`
}

// KmsGrantConstraints represents the KmsGrantConstraints schema from the OpenAPI specification
type KmsGrantConstraints struct {
	Encryptioncontextsubset interface{} `json:"encryptionContextSubset,omitempty"`
	Encryptioncontextequals interface{} `json:"encryptionContextEquals,omitempty"`
}

// ListAnalyzersResponse represents the ListAnalyzersResponse schema from the OpenAPI specification
type ListAnalyzersResponse struct {
	Analyzers interface{} `json:"analyzers"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// InlineArchiveRule represents the InlineArchiveRule schema from the OpenAPI specification
type InlineArchiveRule struct {
	Rulename interface{} `json:"ruleName"`
	Filter interface{} `json:"filter"`
}

// RdsDbSnapshotAttributesMap represents the RdsDbSnapshotAttributesMap schema from the OpenAPI specification
type RdsDbSnapshotAttributesMap struct {
}

// PolicyGeneration represents the PolicyGeneration schema from the OpenAPI specification
type PolicyGeneration struct {
	Completedon interface{} `json:"completedOn,omitempty"`
	Jobid interface{} `json:"jobId"`
	Principalarn interface{} `json:"principalArn"`
	Startedon interface{} `json:"startedOn"`
	Status interface{} `json:"status"`
}

// ListAccessPreviewFindingsRequest represents the ListAccessPreviewFindingsRequest schema from the OpenAPI specification
type ListAccessPreviewFindingsRequest struct {
	Analyzerarn interface{} `json:"analyzerArn"`
	Filter interface{} `json:"filter,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// GetAnalyzedResourceResponseresource represents the GetAnalyzedResourceResponseresource schema from the OpenAPI specification
type GetAnalyzedResourceResponseresource struct {
	Analyzedat interface{} `json:"analyzedAt"`
	ErrorField interface{} `json:"error,omitempty"`
	Ispublic interface{} `json:"isPublic"`
	Resourcearn interface{} `json:"resourceArn"`
	Resourcetype interface{} `json:"resourceType"`
	Sharedvia interface{} `json:"sharedVia,omitempty"`
	Resourceowneraccount interface{} `json:"resourceOwnerAccount"`
	Createdat interface{} `json:"createdAt"`
	Status interface{} `json:"status,omitempty"`
	Updatedat interface{} `json:"updatedAt"`
	Actions interface{} `json:"actions,omitempty"`
}

// GetAnalyzerResponse represents the GetAnalyzerResponse schema from the OpenAPI specification
type GetAnalyzerResponse struct {
	Analyzer GetAnalyzerResponseanalyzer `json:"analyzer"`
}

// Spanstart represents the Spanstart schema from the OpenAPI specification
type Spanstart struct {
	Offset interface{} `json:"offset"`
	Column interface{} `json:"column"`
	Line interface{} `json:"line"`
}

// FindingSummary represents the FindingSummary schema from the OpenAPI specification
type FindingSummary struct {
	Createdat interface{} `json:"createdAt"`
	Resourceowneraccount interface{} `json:"resourceOwnerAccount"`
	Analyzedat interface{} `json:"analyzedAt"`
	Condition interface{} `json:"condition"`
	Ispublic interface{} `json:"isPublic,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
	Status interface{} `json:"status"`
	Id interface{} `json:"id"`
	Principal interface{} `json:"principal,omitempty"`
	Resource interface{} `json:"resource,omitempty"`
	Resourcetype interface{} `json:"resourceType"`
	ErrorField interface{} `json:"error,omitempty"`
	Updatedat interface{} `json:"updatedAt"`
	Action interface{} `json:"action,omitempty"`
}

// ListAccessPreviewsResponse represents the ListAccessPreviewsResponse schema from the OpenAPI specification
type ListAccessPreviewsResponse struct {
	Accesspreviews interface{} `json:"accessPreviews"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListPolicyGenerationsResponse represents the ListPolicyGenerationsResponse schema from the OpenAPI specification
type ListPolicyGenerationsResponse struct {
	Policygenerations interface{} `json:"policyGenerations"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// PathElementsubstring represents the PathElementsubstring schema from the OpenAPI specification
type PathElementsubstring struct {
	Length interface{} `json:"length"`
	Start interface{} `json:"start"`
}

// Trail represents the Trail schema from the OpenAPI specification
type Trail struct {
	Cloudtrailarn interface{} `json:"cloudTrailArn"`
	Regions interface{} `json:"regions,omitempty"`
	Allregions interface{} `json:"allRegions,omitempty"`
}

// FindingSource represents the FindingSource schema from the OpenAPI specification
type FindingSource struct {
	Detail FindingSourcedetail `json:"detail,omitempty"`
	TypeField interface{} `json:"type"`
}

// GetGeneratedPolicyResponse represents the GetGeneratedPolicyResponse schema from the OpenAPI specification
type GetGeneratedPolicyResponse struct {
	Jobdetails GetGeneratedPolicyResponsejobDetails `json:"jobDetails"`
	Generatedpolicyresult GetGeneratedPolicyResponsegeneratedPolicyResult `json:"generatedPolicyResult"`
}

// PathElement represents the PathElement schema from the OpenAPI specification
type PathElement struct {
	Substring PathElementsubstring `json:"substring,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Index interface{} `json:"index,omitempty"`
	Key interface{} `json:"key,omitempty"`
}

// ListAnalyzedResourcesResponse represents the ListAnalyzedResourcesResponse schema from the OpenAPI specification
type ListAnalyzedResourcesResponse struct {
	Analyzedresources interface{} `json:"analyzedResources"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// JobError represents the JobError schema from the OpenAPI specification
type JobError struct {
	Code interface{} `json:"code"`
	Message interface{} `json:"message"`
}

// KmsGrantConfigurationconstraints represents the KmsGrantConfigurationconstraints schema from the OpenAPI specification
type KmsGrantConfigurationconstraints struct {
	Encryptioncontextequals interface{} `json:"encryptionContextEquals,omitempty"`
	Encryptioncontextsubset interface{} `json:"encryptionContextSubset,omitempty"`
}

// S3AccessPointConfigurationnetworkOrigin represents the S3AccessPointConfigurationnetworkOrigin schema from the OpenAPI specification
type S3AccessPointConfigurationnetworkOrigin struct {
	Vpcconfiguration VpcConfiguration `json:"vpcConfiguration,omitempty"` // The proposed virtual private cloud (VPC) configuration for the Amazon S3 access point. VPC configuration does not apply to multi-region access points. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_VpcConfiguration.html">VpcConfiguration</a>.
	Internetconfiguration interface{} `json:"internetConfiguration,omitempty"`
}

// S3BucketConfigurationbucketPublicAccessBlock represents the S3BucketConfigurationbucketPublicAccessBlock schema from the OpenAPI specification
type S3BucketConfigurationbucketPublicAccessBlock struct {
	Ignorepublicacls interface{} `json:"ignorePublicAcls"`
	Restrictpublicbuckets interface{} `json:"restrictPublicBuckets"`
}

// ListFindingsrequest represents the ListFindingsrequest schema from the OpenAPI specification
type ListFindingsrequest struct {
	Nexttoken string `json:"nextToken,omitempty"` // A token used for pagination of results returned.
	Sort ListFindingsrequestsort `json:"sort,omitempty"` // The criteria used to sort.
	Analyzerarn string `json:"analyzerArn"` // The <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources">ARN of the analyzer</a> to retrieve findings from.
	Filter map[string]interface{} `json:"filter,omitempty"` // A filter to match for the findings to return.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results to return in the response.
}

// GetFindingRequest represents the GetFindingRequest schema from the OpenAPI specification
type GetFindingRequest struct {
}

// CreateAccessPreviewrequest represents the CreateAccessPreviewrequest schema from the OpenAPI specification
type CreateAccessPreviewrequest struct {
	Clienttoken string `json:"clientToken,omitempty"` // A client token.
	Configurations map[string]interface{} `json:"configurations"` // Access control configuration for your resource that is used to generate the access preview. The access preview includes findings for external access allowed to the resource with the proposed access control configuration. The configuration must contain exactly one element.
	Analyzerarn string `json:"analyzerArn"` // The <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources">ARN of the account analyzer</a> used to generate the access preview. You can only create an access preview for analyzers with an <code>Account</code> type and <code>Active</code> status.
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// PolicyGenerationDetails represents the PolicyGenerationDetails schema from the OpenAPI specification
type PolicyGenerationDetails struct {
	Principalarn interface{} `json:"principalArn"`
}

// DeleteArchiveRuleRequest represents the DeleteArchiveRuleRequest schema from the OpenAPI specification
type DeleteArchiveRuleRequest struct {
}

// GetAnalyzedResourceRequest represents the GetAnalyzedResourceRequest schema from the OpenAPI specification
type GetAnalyzedResourceRequest struct {
}

// ListFindingsRequestsort represents the ListFindingsRequestsort schema from the OpenAPI specification
type ListFindingsRequestsort struct {
	Attributename interface{} `json:"attributeName,omitempty"`
	Orderby interface{} `json:"orderBy,omitempty"`
}

// Finding represents the Finding schema from the OpenAPI specification
type Finding struct {
	Resourceowneraccount interface{} `json:"resourceOwnerAccount"`
	Sources interface{} `json:"sources,omitempty"`
	Analyzedat interface{} `json:"analyzedAt"`
	Createdat interface{} `json:"createdAt"`
	ErrorField interface{} `json:"error,omitempty"`
	Resourcetype interface{} `json:"resourceType"`
	Status interface{} `json:"status"`
	Ispublic interface{} `json:"isPublic,omitempty"`
	Principal interface{} `json:"principal,omitempty"`
	Updatedat interface{} `json:"updatedAt"`
	Action interface{} `json:"action,omitempty"`
	Id interface{} `json:"id"`
	Resource interface{} `json:"resource,omitempty"`
	Condition interface{} `json:"condition"`
}

// Configuration represents the Configuration schema from the OpenAPI specification
type Configuration struct {
	S3bucket Configurations3Bucket `json:"s3Bucket,omitempty"`
	Snstopic ConfigurationsnsTopic `json:"snsTopic,omitempty"`
	Sqsqueue ConfigurationsqsQueue `json:"sqsQueue,omitempty"`
	Ecrrepository ConfigurationecrRepository `json:"ecrRepository,omitempty"`
	Kmskey ConfigurationkmsKey `json:"kmsKey,omitempty"`
	Iamrole ConfigurationiamRole `json:"iamRole,omitempty"`
	Rdsdbsnapshot ConfigurationrdsDbSnapshot `json:"rdsDbSnapshot,omitempty"`
	Ebssnapshot ConfigurationebsSnapshot `json:"ebsSnapshot,omitempty"`
	Secretsmanagersecret ConfigurationsecretsManagerSecret `json:"secretsManagerSecret,omitempty"`
	Efsfilesystem ConfigurationefsFileSystem `json:"efsFileSystem,omitempty"`
	Rdsdbclustersnapshot ConfigurationrdsDbClusterSnapshot `json:"rdsDbClusterSnapshot,omitempty"`
}

// ConfigurationkmsKey represents the ConfigurationkmsKey schema from the OpenAPI specification
type ConfigurationkmsKey struct {
	Grants interface{} `json:"grants,omitempty"`
	Keypolicies interface{} `json:"keyPolicies,omitempty"`
}

// ConfigurationrdsDbSnapshot represents the ConfigurationrdsDbSnapshot schema from the OpenAPI specification
type ConfigurationrdsDbSnapshot struct {
	Attributes interface{} `json:"attributes,omitempty"`
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
}

// SecretsManagerSecretConfiguration represents the SecretsManagerSecretConfiguration schema from the OpenAPI specification
type SecretsManagerSecretConfiguration struct {
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
	Secretpolicy interface{} `json:"secretPolicy,omitempty"`
}

// S3BucketConfiguration represents the S3BucketConfiguration schema from the OpenAPI specification
type S3BucketConfiguration struct {
	Bucketpolicy interface{} `json:"bucketPolicy,omitempty"`
	Bucketpublicaccessblock S3BucketConfigurationbucketPublicAccessBlock `json:"bucketPublicAccessBlock,omitempty"`
	Accesspoints interface{} `json:"accessPoints,omitempty"`
	Bucketaclgrants interface{} `json:"bucketAclGrants,omitempty"`
}

// GeneratedPolicyPropertiescloudTrailProperties represents the GeneratedPolicyPropertiescloudTrailProperties schema from the OpenAPI specification
type GeneratedPolicyPropertiescloudTrailProperties struct {
	Trailproperties interface{} `json:"trailProperties"`
	Endtime interface{} `json:"endTime"`
	Starttime interface{} `json:"startTime"`
}

// S3PublicAccessBlockConfiguration represents the S3PublicAccessBlockConfiguration schema from the OpenAPI specification
type S3PublicAccessBlockConfiguration struct {
	Ignorepublicacls interface{} `json:"ignorePublicAcls"`
	Restrictpublicbuckets interface{} `json:"restrictPublicBuckets"`
}

// GetAnalyzerRequest represents the GetAnalyzerRequest schema from the OpenAPI specification
type GetAnalyzerRequest struct {
}

// ListAccessPreviewsRequest represents the ListAccessPreviewsRequest schema from the OpenAPI specification
type ListAccessPreviewsRequest struct {
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Path interface{} `json:"path"`
	Span Locationspan `json:"span"`
}

// AnalyzedResource represents the AnalyzedResource schema from the OpenAPI specification
type AnalyzedResource struct {
	Resourceowneraccount interface{} `json:"resourceOwnerAccount"`
	Createdat interface{} `json:"createdAt"`
	Status interface{} `json:"status,omitempty"`
	Updatedat interface{} `json:"updatedAt"`
	Actions interface{} `json:"actions,omitempty"`
	Analyzedat interface{} `json:"analyzedAt"`
	ErrorField interface{} `json:"error,omitempty"`
	Ispublic interface{} `json:"isPublic"`
	Resourcearn interface{} `json:"resourceArn"`
	Resourcetype interface{} `json:"resourceType"`
	Sharedvia interface{} `json:"sharedVia,omitempty"`
}

// ConfigurationrdsDbClusterSnapshot represents the ConfigurationrdsDbClusterSnapshot schema from the OpenAPI specification
type ConfigurationrdsDbClusterSnapshot struct {
	Kmskeyid interface{} `json:"kmsKeyId,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// AclGrantee represents the AclGrantee schema from the OpenAPI specification
type AclGrantee struct {
	Id interface{} `json:"id,omitempty"`
	Uri interface{} `json:"uri,omitempty"`
}

// ListAnalyzersRequest represents the ListAnalyzersRequest schema from the OpenAPI specification
type ListAnalyzersRequest struct {
}
