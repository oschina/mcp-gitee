package pulls

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ApprovePullReviewToolName = "approve_pull_review"
	CancelPullReviewToolName  = "cancel_pull_review"
)

var ApprovePullReviewTool = func() mcp.Tool {
	options := utils.CombineOptions(
		BasicOptions,
		[]mcp.ToolOption{
			mcp.WithDescription("Approve a pull request review"),
			mcp.WithNumber(
				"number",
				mcp.Description("The number of the pull request, must be an integer, not a float"),
				mcp.Required(),
			),
			mcp.WithBoolean(
				"force",
				mcp.Description("Whether to force approve the pull request, only available for administrators"),
				mcp.DefaultBool(false),
			),
		},
	)
	return mcp.NewTool(ApprovePullReviewToolName, options...)
}()

var CancelPullReviewTool = func() mcp.Tool {
	options := utils.CombineOptions(
		BasicOptions,
		[]mcp.ToolOption{
			mcp.WithDescription("Reset the review status of a pull request"),
			mcp.WithNumber(
				"number",
				mcp.Description("The number of the pull request, must be an integer, not a float"),
				mcp.Required(),
			),
			mcp.WithBoolean(
				"reset_all",
				mcp.Description("Whether to reset all reviewers, only available for administrators"),
				mcp.DefaultBool(false),
			),
		},
	)
	return mcp.NewTool(CancelPullReviewToolName, options...)
}()

func ApprovePullReviewHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)

	if checkResult, err := utils.CheckRequired(args, "owner", "repo", "number"); err != nil {
		return checkResult, err
	}

	owner := args["owner"].(string)
	repo := args["repo"].(string)

	number, err := utils.SafelyConvertToInt(args["number"])
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	apiUrl := fmt.Sprintf("/repos/%s/%s/pulls/%d/review", owner, repo, number)
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))
	return giteeClient.HandleMCPResult(nil)
}

func CancelPullReviewHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)

	if checkResult, err := utils.CheckRequired(args, "owner", "repo", "number"); err != nil {
		return checkResult, err
	}

	owner := args["owner"].(string)
	repo := args["repo"].(string)

	number, err := utils.SafelyConvertToInt(args["number"])
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	apiUrl := fmt.Sprintf("/repos/%s/%s/pulls/%d/assignees", owner, repo, number)
	giteeClient := utils.NewGiteeClient("PATCH", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))
	return giteeClient.HandleMCPResult(nil)
}
