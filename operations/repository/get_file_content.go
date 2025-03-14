package repository

import (
	"context"
	"fmt"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"net/url"
)

const (
	// GetFileContentToolName is the tool name for getting file content
	GetFileContentToolName = "get_file_content"
)

var GetFileContentTool = mcp.NewTool(
	GetFileContentToolName,
	mcp.WithDescription("Get the content of the specified file in the repository"),
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
		"path",
		mcp.Description("The path of the file"),
		mcp.Required(),
	),
	mcp.WithString(
		"ref",
		mcp.Description("The branch name or commit ID"),
	),
)

func GetFileContentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner := request.Params.Arguments["owner"].(string)
	repo := request.Params.Arguments["repo"].(string)
	path := request.Params.Arguments["path"].(string)
	ref, ok := request.Params.Arguments["ref"].(string)
	if !ok {
		ref = ""
	}
	apiUrl := fmt.Sprintf("/repos/%s/%s/contents/%s", owner, repo, url.QueryEscape(path))
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithQuery(map[string]interface{}{"ref": ref}))
	var fileContents interface{}
	return giteeClient.HandleMCPResult(&fileContents)
}
