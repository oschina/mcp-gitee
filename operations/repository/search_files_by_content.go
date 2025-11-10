package repository

import (
	"context"

	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	SearchFilesByContentToolName = "search_files_by_content"
)

var SearchFilesByContentTool = mcp.NewTool(
	SearchFilesByContentToolName,
	mcp.WithDescription("Search files by content in a repository"),
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
	mcp.WithString(
		"query",
		mcp.Description("The search keywords"),
		mcp.Required(),
	),
	mcp.WithString(
		"ref",
		mcp.Description("Branch, tag, or commit to search; defaults to the repository's default branch"),
	),
	mcp.WithNumber(
		"limit",
		mcp.Description("Maximum number of search results to return (1-100)"),
		mcp.DefaultNumber(20),
	),
	mcp.WithString(
		"paths",
		mcp.Description("Comma-separated list of paths or glob patterns to limit the search scope"),
	),
	mcp.WithNumber(
		"before_context",
		mcp.Description("Number of context lines to include before each match"),
	),
	mcp.WithNumber(
		"after_context",
		mcp.Description("Number of context lines to include after each match"),
	),
	mcp.WithBoolean(
		"literal_pathspec",
		mcp.Description("Treat provided paths as literal strings and disable glob matching"),
		mcp.DefaultBool(false),
	),
)

func SearchFilesByContentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)

	if checkResult, err := utils.CheckRequired(args, "owner", "repo", "query"); err != nil {
		return checkResult, err
	}

	owner := args["owner"].(string)
	repo := args["repo"].(string)

	apiUrl := "/repos/" + owner + "/" + repo + "/search/files_by_content"
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(args))

	var result []string
	return giteeClient.HandleMCPResult(&result)
}
