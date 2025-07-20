# Gitee MCP Server

Gitee MCP Server is a Model Context Protocol (MCP) server implementation for Gitee. It provides a set of tools for interacting with Gitee's API, allowing AI assistants to manage repositories, issues, pull requests, and more.

[![Install MCP Server](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=gitee&config=eyJjb21tYW5kIjoibnB4IC15IEBnaXRlZS9tY3AtZ2l0ZWVAbGF0ZXN0IiwiZW52Ijp7IkdJVEVFX0FDQ0VTU19UT0tFTiI6Ijx5b3VyIHBlcnNvbmFsIGFjY2VzcyB0b2tlbj4ifX0%3D)

## Features

- Interact with Gitee repositories, issues, pull requests, and notifications
- Configurable API base URL to support different Gitee instances
- Command-line flags for easy configuration
- Supports both personal, organization, and enterprise operations
- Dynamic toolset enable/disable

<details>
<summary><b>Practical scenario: Obtain Issue from the repository, implement and create a Pull Request</b></summary>

1. Get repository Issues
![get_repo_issues](./docs/images/get_repo_issues.jpg)
2. Implement coding & create Pull Request based on Issue details
![implement_issue](./docs/images/implement_issue.jpg)
3. Comment & Close Issue
![comment_and_close_issue](./docs/images/comment_and_close_issue.jpg)
</details>

## Installation(This step can be skipped directly when starting npx)

### Prerequisites

- Go 1.23.0 or higher
- Gitee account with an access token, [Go to get](https://gitee.com/profile/personal_access_tokens)

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://gitee.com/oschina/mcp-gitee.git
   cd mcp-gitee
   ```

2. Build the project:
   ```bash
   make build
   ```
   Move ./bin/mcp-gitee PATH env

### Use go install
   ```bash
   go install gitee.com/oschina/mcp-gitee@latest
   ```

## Usage

Check mcp-gitee version:

```bash
mcp-gitee --version
```

## MCP Hosts Configuration
<div align="center">
  <a href="docs/install/claude.md" title="Claude"><img src="docs/install/logos/Claude.png" width=80 height=80></a>
  <a href="docs/install/cursor.md" title="Cursor"><img src="docs/install/logos/Cursor.png" width=80 height=80></a>
  <a href="docs/install/trae.md" title="Trae"><img src="docs/install/logos/Trae.png" width=80 height=80></a>
  <a href="docs/install/cline.md" title="Cline"><img src="docs/install/logos/Cline.png" width=80 height=80></a>
  <a href="docs/install/windsurf.md" title="Windsurf"><img src="docs/install/logos/Windsurf.png" width=80 height=80></a>
</div>

config example: [Click to view more application configuration](./docs/install/)
- Connect to the official remote mcp-gitee server (no installation required)
```json
{
  "mcpServers": {
    "gitee": {
      "url": "https://api.gitee.com/mcp",
      "headers": {
        "Authorization": "Bearer <your personal access token>"
      }
    }
  }
}
```

- npx
```json
{
  "mcpServers": {
    "gitee": {
      "command": "npx",
      "args": [
        "-y",
        "@gitee/mcp-gitee@latest"
      ],
      "env": {
        "GITEE_API_BASE": "https://gitee.com/api/v5",
        "GITEE_ACCESS_TOKEN": "<your personal access token>"
      }
    }
  }
}
```
- executable
```json
{
  "mcpServers": {
    "gitee": {
      "command": "mcp-gitee",
      "env": {
        "GITEE_API_BASE": "https://gitee.com/api/v5",
        "GITEE_ACCESS_TOKEN": "<your personal access token>"
      }
    }
  }
}
```

### Command-line Options

- `--token`: Gitee access token
- `--api-base`: Gitee API base URL (default: https://gitee.com/api/v5)
- `--version`: Show version information
- `--transport`: Transport type (stdio、sse or http, default: stdio)
- `--address`: The host and port to start the server on (default: localhost:8000)
- `--enabled-toolsets`: Comma-separated list of tools to enable (if specified, only these tools will be enabled)
- `--disabled-toolsets`: Comma-separated list of tools to disable

### Environment Variables

You can also configure the server using environment variables:

- `GITEE_ACCESS_TOKEN`: Gitee access token
- `GITEE_API_BASE`: Gitee API base URL
- `ENABLED_TOOLSETS`: Comma-separated list of tools to enable
- `DISABLED_TOOLSETS`: Comma-separated list of tools to disable

### Toolset Management

Toolset management supports two modes:

1. Enable specified tools (whitelist mode):
   - Use `--enabled-toolsets` parameter or `ENABLED_TOOLSETS` environment variable
   - Specify after, only listed tools will be enabled, others will be disabled
   - Example: `--enabled-toolsets="list_user_repos,get_file_content"`

2. Disable specified tools (blacklist mode):
   - Use `--disabled-toolsets` parameter or `DISABLED_TOOLSETS` environment variable
   - Specify after, listed tools will be disabled, others will be enabled
   - Example: `--disabled-toolsets="list_user_repos,get_file_content"`

Note:
- If both `enabled-toolsets` and `disabled-toolsets` are specified, `enabled-toolsets` takes precedence
- Tool names are case-sensitive

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Available Tools

The server provides various tools for interacting with Gitee:

| Tool                                | Category | Description |
|-------------------------------------|----------|-------------|
| **list_user_repos**                 | Repository | List user authorized repositories |
| **get_file_content**                | Repository | Get the content of a file in a repository |
| **create_user_repo**                | Repository | Create a user repository |
| **create_org_repo**                 | Repository | Create an organization repository |
| **create_enter_repo**               | Repository | Create an enterprise repository |
| **fork_repository**                 | Repository | Fork a repository |
| **create_release**                  | Repository | Create a release for a repository |
| **list_releases**                   | Repository | List repository releases |
| **search_open_source_repositories** | Repository | Search open source repositories on Gitee |
| **list_repo_pulls**                 | Pull Request | List pull requests in a repository |
| **merge_pull**                      | Pull Request | Merge a pull request |
| **create_pull**                     | Pull Request | Create a pull request |
| **update_pull**                     | Pull Request | Update a pull request |
| **get_pull_detail**                 | Pull Request | Get details of a pull request |
| **comment_pull**                    | Pull Request | Comment on a pull request |
| **list_pull_comments**              | Pull Request | List all comments for a pull request |
| **get_diff_files**                  | Pull Request | Get a pull request diff files |
| **create_issue**                    | Issue | Create an issue |
| **update_issue**                    | Issue | Update an issue |
| **get_repo_issue_detail**           | Issue | Get details of a repository issue |
| **list_repo_issues**                | Issue | List repository issues |
| **comment_issue**                   | Issue | Comment on an issue |
| **list_issue_comments**             | Issue | List comments on an issue |
| **get_user_info**                   | User | Get current authenticated user information |
| **search_users**                    | User | Search for users |
| **list_user_notifications**         | Notification | List user notifications |

## Contribution

We welcome contributions from the open-source community! If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure the code is well-documented.
4. Submit a pull request with a clear description of your changes.

For more information, please refer to the [CONTRIBUTING](CONTRIBUTING.md) file.
