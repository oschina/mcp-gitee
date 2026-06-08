# opencode

Add the remote MCP server to `~/.config/opencode/opencode.json`:

```json
{
  "mcp": {
    "gitee": {
      "type": "remote",
      "url": "https://api.gitee.com/mcp",
      "headers": {
        "Authorization": "Bearer <your personal access token>"
      },
      "enabled": true
    }
  }
}
```
