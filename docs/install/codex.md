# Codex

Add the remote MCP server to `~/.codex/config.toml`:

```toml
[mcp_servers.gitee]
url = "https://api.gitee.com/mcp"
bearer_token_env_var = "GITEE_ACCESS_TOKEN"
```

Then set `GITEE_ACCESS_TOKEN` in your shell environment.

You can also add it with the Codex CLI:

```bash
codex mcp add gitee --url https://api.gitee.com/mcp --bearer-token-env-var GITEE_ACCESS_TOKEN
```
