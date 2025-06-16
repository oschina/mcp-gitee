package issues

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	// Tool name constants
	GetRepoIssueDetailToolName = "get_repo_issue_detail"
)

// Basic issue get options
var basicIssueGetOptions = []mcp.ToolOption{
	mcp.WithString(
		"number",
		mcp.Description("Issue number (case sensitive, no # prefix needed)"),
		mcp.Required(),
	),
}

// Repository issue get options
var repoIssueGetOptions = []mcp.ToolOption{
	mcp.WithString(
		"owner",
		mcp.Description("The space address to which the repository belongs (the address path of the enterprise, organization or individual)"),
		mcp.Required(),
	),
	mcp.WithString(
		"repo",
		mcp.Description("The path of the repository"),
		mcp.Required(),
	),
}

// Enterprise issue get options
var enterpriseIssueGetOptions = []mcp.ToolOption{
	mcp.WithString(
		"enterprise",
		mcp.Description("Enterprise path"),
		mcp.Required(),
	),
}

// Endpoint configurations
var getIssueConfigs = map[string]types.EndpointConfig{
	GetRepoIssueDetailToolName: {
		UrlTemplate: "/repos/%s/%s/issues/%s",
		PathParams:  []string{"owner", "repo", "number"},
	},
}

// NewGetIssueDetailTool creates a tool for getting issue details
func NewGetIssueDetailTool(getType string) mcp.Tool {
	_, ok := getIssueConfigs[getType]
	if !ok {
		panic("invalid get issue type: " + getType)
	}

	var options []mcp.ToolOption
	options = append(options, basicIssueGetOptions...)

	switch getType {
	case GetRepoIssueDetailToolName:
		options = append(options, mcp.WithDescription("Get the detail of an issue"))
		options = append(options, repoIssueGetOptions...)
	}

	return mcp.NewTool(getType, options...)
}

// GetIssueDetailHandleFunc handles the request to get issue details
func GetIssueDetailHandleFunc(getType string) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		config, ok := getIssueConfigs[getType]
		if !ok {
			errMsg := fmt.Sprintf("unsupported get issue type: %s", getType)
			return mcp.NewToolResultError(errMsg), fmt.Errorf(errMsg)
		}

		args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
		apiUrl := config.UrlTemplate
		if len(config.PathParams) > 0 {
			apiUrlArgs := make([]interface{}, len(config.PathParams))
			for idx, param := range config.PathParams {
				value, ok := args[param].(string)
				if !ok {
					errMsg := fmt.Sprintf("missing required path parameter: %s", param)
					return mcp.NewToolResultError(errMsg), fmt.Errorf(errMsg)
				}
				apiUrlArgs[idx] = value
			}
			apiUrl = fmt.Sprintf(apiUrl, apiUrlArgs...)
		}

		// Handle optional query parameters
		queryParams := make(map[string]interface{})
		if accessToken, ok := args["access_token"]; ok && accessToken != "" {
			queryParams["access_token"] = accessToken
		}

		giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(queryParams))

		// 使用 HandleMCPResult 方法简化处理逻辑
		issue := &types.BasicIssue{}
		return giteeClient.HandleMCPResult(issue)
	}
}
