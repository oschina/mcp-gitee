package repository

import (
	"context"
	"fmt"

	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateRepoToolName = "create_repo"
)

var createRepoConfigs = map[string]types.EndpointConfig{
	"user": {
		UrlTemplate: "/user/repos",
	},
	"org": {
		UrlTemplate: "/orgs/%s/repos",
		PathParams:  []string{"org"},
	},
	"enterprise": {
		UrlTemplate: "/enterprises/%s/repos",
		PathParams:  []string{"enterprise"},
	},
}

var CreateRepoTool = mcp.NewTool(
	CreateRepoToolName,
	mcp.WithDescription("Create a repository (user, org, or enterprise)"),
	mcp.WithString(
		"owner_type",
		mcp.Description("Owner type: user (personal), org (organization), or enterprise"),
		mcp.Required(),
		mcp.Enum("user", "org", "enterprise"),
	),
	mcp.WithString(
		"name",
		mcp.Description("Repository name"),
		mcp.Required(),
	),
	mcp.WithString(
		"description",
		mcp.Description("Repository description"),
	),
	mcp.WithString(
		"homepage",
		mcp.Description("Repository homepage"),
	),
	mcp.WithBoolean(
		"auto_init",
		mcp.Description("Whether to initialize the repository with a README file"),
		mcp.DefaultBool(false),
	),
	mcp.WithBoolean(
		"private",
		mcp.Description("Whether the repository is private"),
		mcp.DefaultBool(true),
	),
	mcp.WithString(
		"path",
		mcp.Description("Repository path (for customization)"),
	),
	mcp.WithString(
		"org",
		mcp.Description("Organization path (required when owner_type is org)"),
	),
	mcp.WithString(
		"enterprise",
		mcp.Description("Enterprise path (required when owner_type is enterprise)"),
	),
)

func CreateRepoHandleFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)

	ownerType, ok := args["owner_type"].(string)
	if !ok {
		return mcp.NewToolResultError("missing required parameter: owner_type"), fmt.Errorf("missing owner_type")
	}

	config, ok := createRepoConfigs[ownerType]
	if !ok {
		return mcp.NewToolResultError("invalid owner_type: must be user, org, or enterprise"), fmt.Errorf("invalid owner_type: %s", ownerType)
	}

	apiUrl := config.UrlTemplate
	if len(config.PathParams) > 0 {
		apiUrlArgs := make([]interface{}, len(config.PathParams))
		for i, param := range config.PathParams {
			value, ok := args[param].(string)
			if !ok || value == "" {
				return mcp.NewToolResultError(fmt.Sprintf("missing required parameter: %s", param)), fmt.Errorf("missing required path parameter: %s", param)
			}
			apiUrlArgs[i] = value
		}
		apiUrl = fmt.Sprintf(apiUrl, apiUrlArgs...)
	}

	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithContext(ctx), utils.WithPayload(args))
	repo := &types.Project{}
	return giteeClient.HandleMCPResult(repo)
}