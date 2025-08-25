package main

import (
	"github.com/access-analyzer/mcp-server/config"
	"github.com/access-analyzer/mcp-server/models"
	tools_tags "github.com/access-analyzer/mcp-server/tools/tags"
	tools_access_preview_analyzerarn "github.com/access-analyzer/mcp-server/tools/access_preview_analyzerarn"
	tools_analyzer "github.com/access-analyzer/mcp-server/tools/analyzer"
	tools_policy "github.com/access-analyzer/mcp-server/tools/policy"
	tools_resource "github.com/access-analyzer/mcp-server/tools/resource"
	tools_access_preview "github.com/access-analyzer/mcp-server/tools/access_preview"
	tools_analyzed_resource "github.com/access-analyzer/mcp-server/tools/analyzed_resource"
	tools_finding "github.com/access-analyzer/mcp-server/tools/finding"
	tools_analyzed_resource_analyzerarn_resourcearn "github.com/access-analyzer/mcp-server/tools/analyzed_resource_analyzerarn_resourcearn"
	tools_archive_rule "github.com/access-analyzer/mcp-server/tools/archive_rule"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tags.CreateUntagresourceTool(cfg),
		tools_access_preview_analyzerarn.CreateListaccesspreviewsTool(cfg),
		tools_analyzer.CreateListanalyzersTool(cfg),
		tools_analyzer.CreateCreateanalyzerTool(cfg),
		tools_policy.CreateStartpolicygenerationTool(cfg),
		tools_policy.CreateListpolicygenerationsTool(cfg),
		tools_resource.CreateStartresourcescanTool(cfg),
		tools_analyzer.CreateDeleteanalyzerTool(cfg),
		tools_analyzer.CreateGetanalyzerTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_access_preview.CreateGetaccesspreviewTool(cfg),
		tools_analyzed_resource.CreateListanalyzedresourcesTool(cfg),
		tools_finding.CreateGetfindingTool(cfg),
		tools_analyzer.CreateListarchiverulesTool(cfg),
		tools_analyzer.CreateCreatearchiveruleTool(cfg),
		tools_policy.CreateGetgeneratedpolicyTool(cfg),
		tools_policy.CreateCancelpolicygenerationTool(cfg),
		tools_analyzed_resource_analyzerarn_resourcearn.CreateGetanalyzedresourceTool(cfg),
		tools_analyzer.CreateUpdatearchiveruleTool(cfg),
		tools_analyzer.CreateDeletearchiveruleTool(cfg),
		tools_analyzer.CreateGetarchiveruleTool(cfg),
		tools_archive_rule.CreateApplyarchiveruleTool(cfg),
		tools_access_preview.CreateCreateaccesspreviewTool(cfg),
		tools_access_preview.CreateListaccesspreviewfindingsTool(cfg),
		tools_finding.CreateListfindingsTool(cfg),
		tools_finding.CreateUpdatefindingsTool(cfg),
		tools_policy.CreateValidatepolicyTool(cfg),
	}
}
