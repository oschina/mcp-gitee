package issues

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// CreateIssueToolName is the name of the tool
	CreateIssueToolName = "create_issue"
)

var CreateIssueTool = func() mcp.Tool {
	options := utils.CombineOptions(
		[]mcp.ToolOption{
			mcp.WithDescription("Create an issue"),
		},
		BasicOptions,
		BasicIssueOptions,
	)
	return mcp.NewTool(CreateIssueToolName, options...)
}()

func CreateIssueHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	owner := args["owner"].(string)
	apiUrl := fmt.Sprintf("/repos/%s/issues", owner)
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))
	issue := &types.BasicIssue{}
	return giteeClient.HandleMCPResult(issue)
}
