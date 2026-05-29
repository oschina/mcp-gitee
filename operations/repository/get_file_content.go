package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"gitee.com/oschina/mcp-gitee/operations/types"
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

	// Do the HTTP request
	_, err := giteeClient.Do()
	if err != nil {
		switch {
		case utils.IsAuthError(err):
			return mcp.NewToolResultError("Authentication failed: Please check your Gitee access token"), err
		case utils.IsNetworkError(err):
			return mcp.NewToolResultError("Network error: Unable to connect to Gitee API"), err
		case utils.IsAPIError(err):
			giteeErr := err.(*utils.GiteeError)
			return mcp.NewToolResultError(fmt.Sprintf("API error (%d): %s", giteeErr.Code, giteeErr.Details)), err
		default:
			return mcp.NewToolResultError(err.Error()), err
		}
	}

	body, err := giteeClient.GetRespBody()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to read response body: %s", err.Error())),
			utils.NewInternalError(errors.New(err.Error()))
	}

	// Detect whether the response is a JSON array (directory listing) or object (single file)
	trimmed := bytes.TrimLeft(body, " \t\r\n")

	if len(trimmed) > 0 && trimmed[0] == '[' {
		// Response is an array (directory listing)
		var fileContents []types.FileContent
		return giteeClient.ProcessResponse(body, &fileContents)
	} else {
		// Response is an object (single file)
		var fileContent types.FileContent
		return giteeClient.ProcessResponse(body, &fileContent)
	}
}
