package repository

import (
	"context"
	"fmt"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateReleaseToolName = "create_release"
)

var CreateReleaseTool = mcp.NewTool(
	CreateReleaseToolName,
	mcp.WithDescription("Create a release"),
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
		"tag_name",
		mcp.Description("The name of the tag"),
		mcp.Required(),
	),
	mcp.WithString(
		"name",
		mcp.Description("The name of the release"),
		mcp.Required(),
	),
	mcp.WithString(
		"body",
		mcp.Description("The description of the release"),
		mcp.Required(),
	),
	mcp.WithBoolean(
		"prerelease",
		mcp.Description("Whether the release is a prerelease"),
		mcp.DefaultBool(false),
	),
	mcp.WithString(
		"target_commitish",
		mcp.Description("The branch name or commit SHA"),
	),
)

func CreateReleaseHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner := request.Params.Arguments["owner"].(string)
	repo := request.Params.Arguments["repo"].(string)

	apiUrl := fmt.Sprintf("/repos/%s/%s/releases", owner, repo)

	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))

	release := &types.Release{}
	return giteeClient.HandleMCPResult(release)
}
