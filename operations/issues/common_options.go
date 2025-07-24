package issues

import "github.com/mark3labs/mcp-go/mcp"

var BasicOptions = []mcp.ToolOption{
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
}

var BasicIssueOptions = []mcp.ToolOption{
	mcp.WithString(
		"title",
		mcp.Description("The title of the issue"),
		mcp.Required(),
	),
	mcp.WithString(
		"body",
		mcp.Description("The description of the issue"),
	),
	mcp.WithString(
		"issue_type",
		mcp.Description("Enterprise custom task type, non-enterprise users must consider it as 'task'"),
	),
	mcp.WithString(
		"assignee",
		mcp.Description("The personal space address of the issue assignee"),
	),
	mcp.WithString(
		"collaborators",
		mcp.Description("The personal space addresses of issue collaborators, separated by commas"),
	),
	mcp.WithString(
		"milestone",
		mcp.Description("The milestone number"),
	),
	mcp.WithString(
		"labels",
		mcp.Description("Comma-separated labels, name requirements are between 2-20 in length and non-special characters. Example: bug,performance"),
	),
	mcp.WithString(
		"program",
		mcp.Description("Project ID"),
	),
	mcp.WithBoolean(
		"security_hole",
		mcp.Description("Set as a private issue (default is false)"),
		mcp.DefaultBool(false),
	),
}
