package pulls

import (
	"context"
	"fmt"
	"strconv"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// ListPullCommentsToolName is the name of the tool
	ListPullCommentsToolName = "list_pull_comments"
)

// ListPullCommentsOptions defines the specific options for listing pull request comments
var ListPullCommentsOptions = []mcp.ToolOption{
	mcp.WithNumber(
		"number",
		mcp.Description("The number of the pull request, must be an integer, not a float"),
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

// ListPullCommentsTool defines the tool for listing pull request comments
var ListPullCommentsTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("List all comments for a pull request"),
		},
		BasicOptions,
		ListPullCommentsOptions,
	)
	return mcp.NewTool(ListPullCommentsToolName, options...)
}()

// ListPullCommentsHandleFunc handles the request to list pull request comments
func ListPullCommentsHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract required parameters from the request
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	owner := args["owner"].(string)
	repo := args["repo"].(string)

	// Extract and convert number parameter
	numberArg, exists := args["number"]
	if !exists {
		return mcp.NewToolResultError("Missing required parameter: number"),
			utils.NewParamError("number", "parameter is required")
	}

	number, err := utils.SafelyConvertToInt(numberArg)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	// Construct the API URL for listing pull request comments
	apiUrl := fmt.Sprintf("/repos/%s/%s/pulls/%d/comments", owner, repo, number)

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
	var comments []types.PullComment

	// Handle the API call and return the result
	return giteeClient.HandleMCPResult(&comments)
}
