package repository

import (
	"context"
	"fmt"
	"gitee.com/oschina/mcp-gitee/operations/types"
	"gitee.com/oschina/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	CreateUserRepo  = "create_user_repo"
	CreateOrgRepo   = "create_org_repo"
	CreateEnterRepo = "create_enterprise_repo"
)

var commonRepoCreateOptions = []mcp.ToolOption{
	mcp.WithString("name", mcp.Description("Repository name"), mcp.Required()),
	mcp.WithString("description", mcp.Description("Repository description")),
	mcp.WithString("homepage", mcp.Description("Repository homepage")),
	mcp.WithBoolean("auto_init", mcp.Description("Whether to initialize the repository with a README file"), mcp.DefaultBool(false)),
	mcp.WithBoolean("private", mcp.Description("Whether the repository is private"), mcp.DefaultBool(true)),
	mcp.WithString("path", mcp.Description("Repository path")),
}

var orgRepoCreateOptions = []mcp.ToolOption{
	mcp.WithString("org", mcp.Description("Org path"), mcp.Required()),
}

var enterpriseCreateOptions = []mcp.ToolOption{
	mcp.WithString("enterprise", mcp.Description("Enterprise path"), mcp.Required()),
}

var createConfigs = map[string]types.EndpointConfig{
	CreateUserRepo: {
		UrlTemplate: "/user/repos",
	},
	CreateOrgRepo: {
		UrlTemplate: "/orgs/%s/repos",
		PathParams:  []string{"org"},
	},
	CreateEnterRepo: {
		UrlTemplate: "/enterprises/%s/repos",
		PathParams:  []string{"enterprise"},
	},
}

func NewCreateRepoTool(createType string) mcp.Tool {
	_, ok := createConfigs[createType]
	if !ok {
		panic("invalid create type: " + createType)
	}

	options := commonRepoCreateOptions
	switch createType {
	case CreateOrgRepo:
		options = append(options, mcp.WithDescription("Create a org repository"))
		options = append(options, orgRepoCreateOptions...)
	case CreateEnterRepo:
		options = append(options, mcp.WithDescription("Create a enterprise repository"))
		options = append(options, enterpriseCreateOptions...)
	default:
		options = append(options, mcp.WithDescription("Create a user repository"))
	}

	return mcp.NewTool(createType, options...)
}

func CreateRepoHandleFunc(createType string) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		config, ok := createConfigs[createType]
		if !ok {
			errMsg := fmt.Sprintf("unsupported create type: %s", createType)
			return mcp.NewToolResultError(errMsg), fmt.Errorf(errMsg)
		}

		apiUrl := config.UrlTemplate
		if len(config.PathParams) > 0 {
			args := make([]interface{}, len(config.PathParams))
			for i, param := range config.PathParams {
				value, ok := request.Params.Arguments[param].(string)
				if !ok {
					return nil, fmt.Errorf("missing required path parameter: %s", param)
				}
				args[i] = value
			}
			apiUrl = fmt.Sprintf(apiUrl, args...)
		}

		giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
		repo := &types.Project{}
		return giteeClient.HandleMCPResult(repo)
	}
}
