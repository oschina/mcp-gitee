package repository

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
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
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	owner := args["owner"].(string)
	repo := args["repo"].(string)
	path := args["path"].(string)
	ref, ok := args["ref"].(string)
	if !ok {
		ref = ""
	}

	// Encode each path segment individually to avoid encoding / to %2F
	parts := strings.Split(path, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	escapedPath := strings.Join(parts, "/")

	apiUrl := fmt.Sprintf("/repos/%s/%s/contents/%s", owner, repo, escapedPath)
	giteeClient := utils.NewGiteeClient("GET", apiUrl, utils.WithContext(ctx), utils.WithQuery(map[string]interface{}{"ref": ref}))

	// Use the specialized handler for polymorphic file content responses
	return giteeClient.HandleMCPResultForFileContent()
}
