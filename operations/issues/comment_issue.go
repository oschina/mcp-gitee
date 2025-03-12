package issues

import (
	"context"
	"fmt"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// CommentIssueToolName is the name of the tool
	CommentIssueToolName = "comment_issue"
)

// CommentIssueOptions defines the specific options for commenting on an issue
var CommentIssueOptions = []mcp.ToolOption{
	mcp.WithString(
		"number",
		mcp.Description("Issue number (case sensitive, no # prefix needed)"),
		mcp.Required(),
	),
	mcp.WithString(
		"body",
		mcp.Description("The contents of the comment"),
		mcp.Required(),
	),
}

// CommentIssueTool defines the tool for commenting on an issue
var CommentIssueTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Create a comment on a repository issue"),
		},
		BasicOptions,
		CommentIssueOptions,
	)
	return mcp.NewTool(CommentIssueToolName, options...)
}()

// CommentIssueHandleFunc handles the request to comment on an issue
func CommentIssueHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract required parameters from the request
	owner := request.Params.Arguments["owner"].(string)
	repo := request.Params.Arguments["repo"].(string)
	number := request.Params.Arguments["number"].(string)

	// Construct the API URL for creating a comment on an issue
	apiUrl := fmt.Sprintf("/repos/%s/%s/issues/%s/comments", owner, repo, number)

	// Create a new Gitee client with the POST method and the constructed API URL
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))

	// Define the response structure
	comment := &types.IssueComment{}

	// Handle the API call and return the result
	return giteeClient.HandleMCPResult(comment)
}
