package pulls

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ManagePullReviewToolName = "manage_pull_review"
)

var ManagePullReviewTool = mcp.NewTool(
	ManagePullReviewToolName,
	mcp.WithDescription("Manage a pull request review (approve or cancel)"),
	mcp.WithString(
		"action",
		mcp.Description("Action to perform: approve (submit approval) or cancel (reset review status)"),
		mcp.Required(),
		mcp.Enum("approve", "cancel"),
	),
	mcp.WithString(
		"owner",
		mcp.Description("The space address to which the repository belongs (enterprise, organization or personal path)"),
		mcp.Required(),
	),
	mcp.WithString(
		"repo",
		mcp.Description("The path of the repository"),
		mcp.Required(),
	),
	mcp.WithNumber(
		"number",
		mcp.Description("The number of the pull request, must be an integer, not a float"),
		mcp.Required(),
	),
	mcp.WithBoolean(
		"force",
		mcp.Description("Whether to force approve (only for administrators, valid when action=approve)"),
		mcp.DefaultBool(false),
	),
	mcp.WithBoolean(
		"reset_all",
		mcp.Description("Whether to reset all reviewers (only for administrators, valid when action=cancel)"),
		mcp.DefaultBool(false),
	),
)

func ManagePullReviewHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)

	action, ok := args["action"].(string)
	if !ok {
		return mcp.NewToolResultError("missing required parameter: action"), fmt.Errorf("missing action")
	}

	if checkResult, err := utils.CheckRequired(args, "owner", "repo", "number"); err != nil {
		return checkResult, err
	}

	owner := args["owner"].(string)
	repo := args["repo"].(string)

	number, err := utils.SafelyConvertToInt(args["number"])
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	var apiUrl string
	var method string

	switch action {
	case "approve":
		apiUrl = fmt.Sprintf("/repos/%s/%s/pulls/%d/review", owner, repo, number)
		method = "POST"
	case "cancel":
		apiUrl = fmt.Sprintf("/repos/%s/%s/pulls/%d/assignees", owner, repo, number)
		method = "PATCH"
	default:
		return mcp.NewToolResultError("invalid action: must be approve or cancel"), fmt.Errorf("invalid action: %s", action)
	}

	giteeClient := utils.NewGiteeClient(method, apiUrl, utils.WithContext(ctx), utils.WithPayload(args))
	return giteeClient.HandleMCPResult(nil)
}