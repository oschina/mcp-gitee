package pulls

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

var BasicPullOptions = []mcp.ToolOption{
	mcp.WithString(
		"body",
		mcp.Description("The body of the pull request"),
	),
	mcp.WithString(
		"milestone_number",
		mcp.Description("The milestone number of the pull request"),
	),
	mcp.WithString(
		"labels",
		mcp.Description("The labels of the pull request. example: bug,performance"),
	),
	mcp.WithString(
		"issue",
		mcp.Description("The issue of the pull request"),
	),
	mcp.WithString(
		"assignees",
		mcp.Description("The assignees of the pull request, example: (username1,username2)"),
	),
	mcp.WithString(
		"testers",
		mcp.Description("The testers of the pull request, example: (username1,username2)"),
	),
	mcp.WithNumber(
		"assignees_number",
		mcp.Description("The min number of assignees need accept of the pull request"),
	),
	mcp.WithNumber(
		"testers_number",
		mcp.Description("The min number of testers need accept of the pull request"),
	),
	mcp.WithBoolean(
		"draft",
		mcp.Description("Whether to set the pull request as draft"),
	),
}
