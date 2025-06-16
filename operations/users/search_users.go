package users

import (
	"context"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	SearchUsers = "search_users"
)

var SearchUsersTool = mcp.NewTool(SearchUsers,
	mcp.WithDescription("Search users on Gitee"),
	mcp.WithString(
		"q",
		mcp.Description("Search keywords"),
		mcp.Required(),
	),
	mcp.WithNumber(
		"page",
		mcp.Description("Page number"),
		mcp.DefaultNumber(1),
	),
	mcp.WithNumber(
		"per_page",
		mcp.Description("Number of results per page"),
		mcp.DefaultNumber(20),
	),
)

func SearchUsersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	if checkResult, err := utils.CheckRequired(args, "q"); err != nil {
		return checkResult, err
	}

	apiUrl := "/search/users"
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(args), utils.WithSkipAuth())

	users := make([]types.BasicUser, 0)
	return giteeClient.HandleMCPResult(&users)
}
