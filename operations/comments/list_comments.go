package comments

import (
	"context"
	"fmt"
	"strconv"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListCommentsToolName = "list_comments"
)

var ListCommentsTool = mcp.NewTool(
	ListCommentsToolName,
	mcp.WithDescription("List all comments for an issue or pull request"),
	mcp.WithString(
		"resource_type",
		mcp.Description("Resource type: issue or pull"),
		mcp.Required(),
		mcp.Enum("issue", "pull"),
	),
	mcp.WithString(
		"owner",
		mcp.Description("The space address to which the repository belongs (enterprise, organization or personal path)"),
		mcp.Required(),
	),
	mcp.WithString(
		"repo",
		mcp.Description("The path of the repository"),
		mcp.Required(),
	),
	mcp.WithString(
		"number",
		mcp.Description("Issue or pull request number (case sensitive, no # prefix needed)"),
		mcp.Required(),
	),
	mcp.WithString(
		"since",
		mcp.Description("Only comments updated at or after this time in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ"),
	),
	mcp.WithNumber(
		"page",
		mcp.Description("Current page number"),
		mcp.DefaultNumber(1),
	),
	mcp.WithNumber(
		"per_page",
		mcp.Description("Number of results per page, maximum 100"),
		mcp.DefaultNumber(20),
	),
	mcp.WithString(
		"order",
		mcp.Description("Sort direction: asc (default) or desc"),
		mcp.DefaultString("asc"),
	),
)

func ListCommentsHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)

	resourceType, ok := args["resource_type"].(string)
	if !ok {
		return mcp.NewToolResultError("missing required parameter: resource_type"), fmt.Errorf("missing resource_type")
	}

	owner := args["owner"].(string)
	repo := args["repo"].(string)
	number := args["number"].(string)

	var apiUrl string
	switch resourceType {
	case "issue":
		apiUrl = fmt.Sprintf("/repos/%s/%s/issues/%s/comments", owner, repo, number)
	case "pull":
		apiUrl = fmt.Sprintf("/repos/%s/%s/pulls/%s/comments", owner, repo, number)
	default:
		return mcp.NewToolResultError("invalid resource_type: must be issue or pull"), fmt.Errorf("invalid resource_type: %s", resourceType)
	}

	query := make(map[string]interface{})
	if since, ok := args["since"].(string); ok && since != "" {
		query["since"] = since
	}
	if page, ok := args["page"].(float64); ok {
		query["page"] = strconv.Itoa(int(page))
	}
	if perPage, ok := args["per_page"].(float64); ok {
		query["per_page"] = strconv.Itoa(int(perPage))
	}
	if order, ok := args["order"].(string); ok && order != "" {
		query["order"] = order
	}

	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(query))

	if resourceType == "issue" {
		comments := make([]types.IssueComment, 0)
		return giteeClient.HandleMCPResult(&comments)
	}
	comments := make([]types.PullComment, 0)
	return giteeClient.HandleMCPResult(&comments)
}