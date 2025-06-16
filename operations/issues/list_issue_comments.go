package issues

import (
	"context"
	"fmt"
	"strconv"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// ListIssueCommentsToolName is the name of the tool
	ListIssueCommentsToolName = "list_issue_comments"
)

// ListIssueCommentsOptions defines the specific options for listing issue comments
var ListIssueCommentsOptions = []mcp.ToolOption{
	mcp.WithString(
		"number",
		mcp.Description("Issue number (case sensitive, no # prefix needed)"),
		mcp.Required(),
	),
	mcp.WithString(
		"since",
		mcp.Description("Only comments updated at or after this time are returned. This is a timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ"),
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
		mcp.Description("Sort direction: asc(default), desc"),
		mcp.DefaultString("asc"),
	),
}

// ListIssueCommentsTool defines the tool for listing issue comments
var ListIssueCommentsTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Get all comments for a repository issue"),
		},
		BasicOptions,
		ListIssueCommentsOptions,
	)
	return mcp.NewTool(ListIssueCommentsToolName, options...)
}()

// ListIssueCommentsHandleFunc handles the request to list issue comments
func ListIssueCommentsHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract required parameters from the request
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	owner := args["owner"].(string)
	repo := args["repo"].(string)
	number := args["number"].(string)

	// Construct the API URL for listing issue comments
	apiUrl := fmt.Sprintf("/repos/%s/%s/issues/%s/comments", owner, repo, number)

	// Prepare query parameters
	query := make(map[string]interface{})

	// Add optional parameters if they exist
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

	// Create a new Gitee client with the GET method and the constructed API URL
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(query))

	// Define the response structure
	var comments []types.IssueComment

	// Handle the API call and return the result
	return giteeClient.HandleMCPResult(&comments)
}
