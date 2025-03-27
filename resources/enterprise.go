package resources

import (
	"context"
	"encoding/json"
	"github.com/mark3labs/mcp-go/mcp"
)

var EnterpriseMetaData = mcp.NewResource(
	"metadata://enterprise",
	"enterprise meta data",
	mcp.WithResourceDescription("Enterprise id/name map"),
	mcp.WithMIMEType("application/json"),
)

func EnterpriseMetaDataHandler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	data := map[int]string{
		1: "oschina",
		2: "hightest",
	}

	_data, _ := json.Marshal(data)

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "metadata://enterprise",
			MIMEType: "application/json",
			Text:     string(_data),
		},
	}, nil
}
