package repository

import (
	"context"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	SearchOpenSourceRepositories = "search_open_source_repositories"
)

var SearchReposTool = mcp.NewTool(SearchOpenSourceRepositories,
	mcp.WithDescription("Search open source repositories on Gitee"),
	mcp.WithString(
		"q",
		mcp.Description("Search keywords"),
		mcp.Required(),
	),
	mcp.WithNumber(
		"from",
		mcp.Description("Search starting position"),
		mcp.DefaultNumber(0),
	),
	mcp.WithNumber(
		"size",
		mcp.Description("Page size"),
		mcp.DefaultNumber(20),
	),
	mcp.WithString(
		"sort_by_f",
		mcp.Description("Sorting field"),
		mcp.Enum("star", "last_push_at"),
	),
)

func SearchOpenSourceReposHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if checkResult, err := utils.CheckRequired(request.Params.Arguments, "q"); err != nil {
		return checkResult, err
	}

	apiUrl := "/search/repos"
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithQuery(request.Params.Arguments), utils.WithSkipAuth())

	data := types.PagedResponse[types.SearchProject]{}
	return giteeClient.HandleMCPResult(&data)
}
