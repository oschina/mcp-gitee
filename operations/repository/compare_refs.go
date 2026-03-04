package repository

import (
	"context"
	"fmt"
	"net/url"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CompareRefsToolName = "compare_branches_tags"
)

var CompareRefsTool = mcp.NewTool(
	CompareRefsToolName,
	mcp.WithDescription("Compare two branches, tags, or commits in a repository"),
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
		"base",
		mcp.Description("The base branch, tag, or commit SHA to compare from"),
		mcp.Required(),
	),
	mcp.WithString(
		"head",
		mcp.Description("The head branch, tag, or commit SHA to compare to"),
		mcp.Required(),
	),
)

func CompareRefsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)

	if checkResult, err := utils.CheckRequired(args, "owner", "repo", "base", "head"); err != nil {
		return checkResult, err
	}

	owner := args["owner"].(string)
	repo := args["repo"].(string)
	base := args["base"].(string)
	head := args["head"].(string)

	compareRef := fmt.Sprintf("%s...%s", url.PathEscape(base), url.PathEscape(head))
	apiUrl := fmt.Sprintf("/repos/%s/%s/compare/%s", owner, repo, compareRef)

	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx))

	var result types.CompareResult
	return giteeClient.HandleMCPResult(&result)
}
