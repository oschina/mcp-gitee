package repository

import (
	"context"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// ListUserReposToolName is the name of the tool
	ListUserReposToolName = "list_user_repos"
)

var ListUserReposTool = mcp.NewTool(
	ListUserReposToolName,
	mcp.WithDescription("List user authorized repositories"),
	mcp.WithString(
		"visibility",
		mcp.Description("Visibility of repository"),
		mcp.Enum("public", "private", "all"),
	),
	mcp.WithString(
		"affiliation",
		mcp.Description("Affiliation between user and repository"),
		mcp.Enum("owner", "collaborator", "organization_member", "enterprise_member", "admin"),
	),
	mcp.WithString(
		"type",
		mcp.Description("Filter user repositories: their creation (owner), personal (personal), their membership (member), public (public), private (private), cannot be used together with visibility or affiliation parameters, otherwise a 422 error will be reported"),
		mcp.Enum("all", "owner", "personal", "member", "public", "private"),
	),
	mcp.WithString(
		"sort",
		mcp.Description("Sorting method: creation time (created), update time (updated), last push time (pushed), warehouse ownership and name (full_name). Default: full_name"),
		mcp.Enum("created", "updated", "pushed", "full_name"),
		mcp.DefaultString("full_name"),
	),
	mcp.WithString(
		"direction",
		mcp.Description("Sorting direction: ascending (asc), descending (desc). Default: asc"),
		mcp.Enum("asc", "desc"),
		mcp.DefaultString("asc"),
	),
	mcp.WithString(
		"q",
		mcp.Description("Search keywords"),
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

func ListUserReposHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := "/user/repos"
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithQuery(request.Params.Arguments))

	repositories := make([]types.Project, 0)
	return giteeClient.HandleMCPResult(&repositories)
}
