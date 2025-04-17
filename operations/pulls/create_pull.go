package pulls

import (
	"context"
	"fmt"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// CreatePullToolName is the name of the tool
	CreatePullToolName = "create_pull"
)

var CreatePullTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Create a pull request"),
		},
		BasicOptions,
		BasicPullOptions,
		[]mcp.ToolOption{
			mcp.WithDescription("Create a pull request"),
			mcp.WithString(
				"title",
				mcp.Description("The title of the pull request"),
				mcp.Required(),
			),
			mcp.WithString(
				"head",
				mcp.Description("The source branch of the pull request, Format: branch (master) or: path_with_namespace:branch (oschina/gitee:master)"),
				mcp.Required(),
			),
			mcp.WithString(
				"base",
				mcp.Description("The target branch of the pull request"),
				mcp.Required(),
			),
		},
	)
	return mcp.NewTool(CreatePullToolName, options...)
}()

func CreatePullHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner := request.Params.Arguments["owner"].(string)
	repo := request.Params.Arguments["repo"].(string)
	apiUrl := fmt.Sprintf("/repos/%s/%s/pulls", owner, repo)
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	pull := &types.BasicPull{}
	return giteeClient.HandleMCPResult(pull)
}
