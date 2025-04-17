package repository

import (
	"context"
	"fmt"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ForkRepositoryToolName = "fork_repository"
)

var ForkRepositoryTool = mcp.NewTool(
	ForkRepositoryToolName,
	mcp.WithDescription("Fork a repository"),
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
		"organization",
		mcp.Description("The full address of the organization space to which the repository belongs, default for current user"),
	),
	mcp.WithString(
		"name",
		mcp.Description("The name of the forked repository, default is the same as the original repository"),
	),
	mcp.WithString(
		"path",
		mcp.Description("The path of the forked repository, default is the same as the original repository"),
	),
)

func ForkRepositoryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if checkResult, err := utils.CheckRequired(request.Params.Arguments, "owner", "repo"); err != nil {
		return checkResult, err
	}
	owner := request.Params.Arguments["owner"].(string)
	repo := request.Params.Arguments["repo"].(string)

	apiUrl := fmt.Sprintf("/repos/%s/%s/forks", owner, repo)
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))

	data := &types.Project{}
	return giteeClient.HandleMCPResult(data)
}
