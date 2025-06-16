package notifications

import (
	"context"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	// ListUserNotifications is the instruction to list user notifications
	ListUserNotifications = "list_user_notifications"
)

var ListUserNotificationsTool = mcp.NewTool(
	ListUserNotifications,
	mcp.WithDescription("List all notifications for authorized user"),
	mcp.WithBoolean(
		"unread",
		mcp.Description("Only list unread notifications"),
		mcp.DefaultBool(false),
	),
	mcp.WithBoolean(
		"participating",
		mcp.Description("Only list notifications where the user is directly participating or mentioned"),
		mcp.DefaultBool(false),
	),
	mcp.WithString(
		"type",
		mcp.Description("Filter notifications of a specified type, all: all, event: event notification, referer: @ notification"),
		mcp.Enum("all", "event", "referer"),
		mcp.DefaultString("all"),
	),
	mcp.WithString(
		"since",
		mcp.Description("Only list notifications updated after the given time, requiring the time format to be ISO 8601"),
	),
	mcp.WithString(
		"before",
		mcp.Description("Only list notifications updated before the given time, requiring the time format to be ISO 8601"),
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

func ListUserNotificationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := "/notifications/threads"
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(args))

	notifications := &types.NotificationResult{}
	return giteeClient.HandleMCPResult(notifications)
}
