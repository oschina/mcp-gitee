package pulls

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// GetPullDetailToolName is the name of the tool
	GetPullDetailToolName = "get_pull_detail"
)

var GetPullDetailTool = func() mcp.Tool {
	options := utils.CombineOptions(
		BasicOptions,
		[]mcp.ToolOption{
			mcp.WithDescription("Get a pull request detail"),
			mcp.WithNumber(
				"number",
				mcp.Description("The number of the pull request, must be an integer, not a float"),
				mcp.Required(),
			),
		},
	)
	return mcp.NewTool(GetPullDetailToolName, options...)
}()

func GetPullDetailHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	owner := args["owner"].(string)
	repo := args["repo"].(string)

	numberArg, exists := args["number"]
	if !exists {
		return mcp.NewToolResultError("Missing required parameter: number"),
			utils.NewParamError("number", "parameter is required")
	}

	number, err := utils.SafelyConvertToInt(numberArg)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	apiUrl := fmt.Sprintf("/repos/%s/%s/pulls/%d", owner, repo, number)
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))
	pull := &types.BasicPull{}
	return giteeClient.HandleMCPResult(pull)
}
