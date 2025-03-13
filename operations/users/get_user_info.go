package users

import (
	"context"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	// GetUserInfoToolName is the instruction to get authenticated user information
	GetUserInfoToolName = "get_user_info"
)

// GetUserInfoTool defines the tool for getting authorized user information
var GetUserInfoTool = mcp.NewTool(
	GetUserInfoToolName,
	mcp.WithDescription("This is a tool from the gitee MCP server.\nGet information about the authenticated user"),
	// No parameters needed for this endpoint as it uses the authenticated user's token
)

// GetUserInfoHandler handles the request to get authorized user information
func GetUserInfoHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client := utils.NewGiteeClient("GET", "/user")

	var user types.BasicUser
	return client.HandleMCPResult(&user)
}

// GetUserInfoHandleFunc returns a server.ToolHandlerFunc for handling get user info requests
func GetUserInfoHandleFunc() server.ToolHandlerFunc {
	return GetUserInfoHandler
}
