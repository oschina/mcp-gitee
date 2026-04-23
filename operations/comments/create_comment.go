package comments

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateCommentToolName = "create_comment"
)

var CreateCommentTool = mcp.NewTool(
	CreateCommentToolName,
	mcp.WithDescription("Create a comment on an issue or pull request"),
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
		"body",
		mcp.Description("The contents of the comment"),
		mcp.Required(),
	),
)

func CreateCommentHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))

	if resourceType == "issue" {
		result := &types.IssueComment{}
		return giteeClient.HandleMCPResult(result)
	}
	result := &types.PullComment{}
	return giteeClient.HandleMCPResult(result)
}