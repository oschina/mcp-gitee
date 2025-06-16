package pulls

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// MergePullToolName is the name of the tool
	MergePullToolName = "merge_pull"
)

var MergePullTool = func() mcp.Tool {
	options := utils.CombineOptions(
		BasicOptions,
		[]mcp.ToolOption{
			mcp.WithDescription("Merge a pull request"),
			mcp.WithNumber(
				"number",
				mcp.Description("The number of the pull request"),
				mcp.Required(),
			),
			mcp.WithString(
				"merge_method",
				mcp.Description("The merge method to use"),
				mcp.Enum("merge", "squash", "rebase"),
				mcp.DefaultString("merge"),
			),
			mcp.WithBoolean(
				"prune_source_branch",
				mcp.Description("Whether to delete the source branch after merging"),
			),
			mcp.WithBoolean(
				"close_related_issue",
				mcp.Description("Whether to close the related issue after merging"),
			),
			mcp.WithString(
				"title",
				mcp.Description("The title of the merge commit"),
			),
			mcp.WithString(
				"description",
				mcp.Description("The description of the merge commit"),
			),
		},
	)
	return mcp.NewTool(MergePullToolName, options...)
}()

func MergePullHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	apiUrl := fmt.Sprintf("/repos/%s/%s/pulls/%d/merge", owner, repo, number)
	giteeClient := utils.NewGiteeClient("PUT", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))
	return giteeClient.HandleMCPResult(nil)
}
