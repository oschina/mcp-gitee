# Claude Code

Add the remote MCP server to the `mcpServers` section:

```json
{
  "mcpServers": {
    "gitee": {
      "type": "http",
      "url": "https://api.gitee.com/mcp",
      "headers": {
        "Authorization": "Bearer <your personal access token>"
      }
    }
  }
}
```
