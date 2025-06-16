package repository

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListReleasesToolName = "list_releases"
)

var ListReleasesTool = mcp.NewTool(
	ListReleasesToolName,
	mcp.WithDescription("List repository releases"),
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
		"direction",
		mcp.Description("Optional. Ascending/descending. Not filled in is ascending"),
		mcp.Enum("asc", "desc"),
	),
)

func ListReleasesHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	owner := args["owner"].(string)
	repo := args["repo"].(string)

	apiUrl := fmt.Sprintf("/repos/%s/%s/releases", owner, repo)

	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))

	releases := make([]types.Release, 0)
	return giteeClient.HandleMCPResult(&releases)
}
