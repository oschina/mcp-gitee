package utils

import "github.com/mark3labs/mcp-go/mcp"

func CombineOptions(options ...[]mcp.ToolOption) []mcp.ToolOption {
	var result []mcp.ToolOption
	for _, option := range options {
		result = append(result, option...)
	}
	return result
}
