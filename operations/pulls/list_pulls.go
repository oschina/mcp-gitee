package pulls

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	ListRepoPullsToolName = "list_repo_pulls"
)

var commonPullsOptions = []mcp.ToolOption{
	mcp.WithString(
		"state",
		mcp.Description("State of the pull request"),
		mcp.Enum("open", "closed", "merged", "all"),
	),
	mcp.WithString(
		"head",
		mcp.Description("Source branch of the PR. Format: branch or username:branch"),
	),
	mcp.WithString(
		"base",
		mcp.Description("Target branch name for the pull request"),
	),
	mcp.WithString(
		"sort",
		mcp.Description("Sort field, default by creation time"),
		mcp.Enum("created", "updated", "popularity", "long-running"),
	),
	mcp.WithString(
		"since",
		mcp.Description("Start update time in ISO 8601 format"),
	),
	mcp.WithString(
		"direction",
		mcp.Description("Ascending/descending order"),
		mcp.Enum("asc", "desc"),
	),
	mcp.WithNumber(
		"milestone_number",
		mcp.Description("Milestone number (ID)"),
	),
	mcp.WithString(
		"labels",
		mcp.Description("Comma-separated labels, e.g.: bug,performance"),
	),
	mcp.WithNumber(
		"page",
		mcp.Description("Current page number"),
		mcp.DefaultNumber(1),
	),
	mcp.WithNumber(
		"per_page",
		mcp.Description("Items per page (max 100)"),
		mcp.DefaultNumber(20),
	),
	mcp.WithString(
		"author",
		mcp.Description("PR creator's username"),
	),
	mcp.WithString(
		"assignee",
		mcp.Description("Reviewer's username"),
	),
	mcp.WithString(
		"tester",
		mcp.Description("Tester's username"),
	),
}

var repoPullsOptions = []mcp.ToolOption{
	mcp.WithString(
		"owner",
		mcp.Description("Repository owner's namespace"),
		mcp.Required(),
	),
	mcp.WithString(
		"repo",
		mcp.Description("Repository namespace"),
		mcp.Required(),
	),
}

var listConfigs = map[string]types.EndpointConfig{
	ListRepoPullsToolName: {
		UrlTemplate: "/repos/%s/%s/pulls",
		PathParams:  []string{"owner", "repo"},
	},
}

func NewListPullsTool(listType string) mcp.Tool {
	_, ok := listConfigs[listType]
	if !ok {
		panic("invalid list type: " + listType)
	}
	options := commonPullsOptions
	switch listType {
	case ListRepoPullsToolName:
		options = append(options, mcp.WithDescription("List repository pulls"))
		options = append(options, repoPullsOptions...)
	}
	return mcp.NewTool(listType, options...)
}

func ListPullsHandleFunc(listType string) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		config, ok := listConfigs[listType]
		if !ok {
			errMsg := fmt.Sprintf("invalid list type: %s", listType)
			return mcp.NewToolResultError(errMsg), fmt.Errorf(errMsg)
		}

		args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
		apiUrl := config.UrlTemplate
		if len(config.PathParams) > 0 {
			apiUrlArgs := make([]interface{}, len(config.PathParams))
			for i, param := range config.PathParams {
				value, ok := args[param].(string)
				if !ok {
					return nil, fmt.Errorf("missing required path parameter: %s", param)
				}
				apiUrlArgs[i] = value
			}
			apiUrl = fmt.Sprintf(apiUrl, apiUrlArgs...)
		}

		giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(args))
		pulls := make([]types.BasicPull, 0)
		return giteeClient.HandleMCPResult(&pulls)
	}
}
