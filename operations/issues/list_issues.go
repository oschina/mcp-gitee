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
	ListRepoIssuesToolName = "list_repo_issues"
)

var commonIssuesOptions = []mcp.ToolOption{
	mcp.WithString(
		"state",
		mcp.Description("Issue state: open, progressing, closed, rejected"),
		mcp.Enum("open", "progressing", "closed", "rejected", "all"),
	),
	mcp.WithString(
		"sort",
		mcp.Description("Sort field: creation time (created), update time (updated). Default: created"),
		mcp.Enum("created", "updated"),
	),
	mcp.WithString(
		"direction",
		mcp.Description("Sort direction: ascending (asc) or descending (desc)"),
		mcp.Enum("asc", "desc"),
	),
	mcp.WithString(
		"since",
		mcp.Description("Start update time in ISO 8601 format"),
	),
	mcp.WithString(
		"schedule",
		mcp.Description("Planned start date in format: 2006-01-02T15:04:05Z"),
	),
	mcp.WithString(
		"deadline",
		mcp.Description("Planned completion date in format: 2006-01-02T15:04:05Z"),
	),
	mcp.WithString(
		"created_at",
		mcp.Description("Issue creation time in format: 2006-01-02T15:04:05Z"),
	),
	mcp.WithString(
		"finished_at",
		mcp.Description("Issue completion time in format: 2006-01-02T15:04:05Z"),
	),
	mcp.WithString(
		"filter",
		mcp.Description("Filter parameter: assigned to authorized user (assigned), created by authorized user (created), all issues involving authorized user (all). Default: assigned"),
		mcp.Enum("assigned", "created", "all"),
	),
	mcp.WithString(
		"labels",
		mcp.Description("Comma-separated labels. Example: bug,performance"),
	),
	mcp.WithNumber(
		"page",
		mcp.Description("Current page number"),
		mcp.DefaultNumber(1),
	),
	mcp.WithNumber(
		"per_page",
		mcp.Description("Number of items per page, maximum 100"),
		mcp.DefaultNumber(20),
	),
}

var repoIssuesOptions = []mcp.ToolOption{
	mcp.WithString(
		"owner",
		mcp.Description("Repository owner's namespace (enterprise, organization or personal path)"),
		mcp.Required(),
	),
	mcp.WithString(
		"repo",
		mcp.Description("Repository path"),
		mcp.Required(),
	),
}

var listConfigs = map[string]types.EndpointConfig{
	ListRepoIssuesToolName: {
		UrlTemplate: "/repos/%s/%s/issues",
		PathParams:  []string{"owner", "repo"},
	},
}

func NewListIssuesTool(listType string) mcp.Tool {
	_, ok := listConfigs[listType]
	if !ok {
		panic("invalid list type: " + listType)
	}
	options := commonIssuesOptions
	switch listType {
	case ListRepoIssuesToolName:
		options = append(options, []mcp.ToolOption{
			mcp.WithDescription("List all issues in a repository"),
			mcp.WithString(
				"q",
				mcp.Description("Search keywords"),
			),
		}...)
		options = append(options, repoIssuesOptions...)
	}
	return mcp.NewTool(listType, options...)
}

func ListIssuesHandleFunc(listType string) server.ToolHandlerFunc {
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
		issues := make([]types.BasicIssue, 0)
		return giteeClient.HandleMCPResult(&issues)
	}
}
