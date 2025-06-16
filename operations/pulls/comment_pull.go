package pulls

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// CommentPullToolName is the name of the tool
	CommentPullToolName = "comment_pull"
)

// CommentPullTool defines the tool for commenting on a pull request
var CommentPullTool = func() mcp.Tool {
	options := utils.CombineOptions(
		BasicOptions,
		[]mcp.ToolOption{
			mcp.WithDescription("Create a comment on a pull request"),
			mcp.WithNumber(
				"number",
				mcp.Description("The number of the pull request, must be an integer, not a float"),
				mcp.Required(),
			),
			mcp.WithString(
				"body",
				mcp.Description("The contents of the comment"),
				mcp.Required(),
			),
		},
	)
	return mcp.NewTool(CommentPullToolName, options...)
}()

// CommentPullHandleFunc handles the comment pull request operation
func CommentPullHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	owner := args["owner"].(string)
	repo := args["repo"].(string)
	body := args["body"].(string)

	numberArg, exists := args["number"]
	if !exists {
		return mcp.NewToolResultError("Missing required parameter: number"),
			utils.NewParamError("number", "parameter is required")
	}

	number, err := utils.SafelyConvertToInt(numberArg)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	// Prepare the API URL for commenting on a pull request
	apiUrl := fmt.Sprintf("/repos/%s/%s/pulls/%d/comments", owner, repo, number)

	// Create payload with the comment body
	payload := map[string]interface{}{
		"body": body,
	}

	// Create a new Gitee client with the POST method and payload
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithContext(ctx), utils.WithPayload(payload))

	// Execute the request and handle the result
	return giteeClient.HandleMCPResult(nil)
}
