# Gitee MCP Server

Gitee MCP 服务器是一个用于 Gitee 的模型上下文协议（Model Context Protocol，MCP）服务器实现。它提供了一系列与 Gitee API 交互的工具，使 AI 助手能够管理仓库、问题、拉取请求等。

## 功能特点

- 与 Gitee 仓库、问题、拉取请求和通知进行交互
- 可配置的 API 基础 URL，支持不同的 Gitee 实例
- 命令行标志，便于配置
- 支持个人、组织和企业操作

<details>
<summary><b>实战场景：从仓库获取 Issue，实现并创建 Pull Request</b></summary>

1. 获取当前仓库 Issues
![get_repo_issues](./docs/images/get_repo_issues.jpg)
2. 根据 Issue 详情实现编码 & 创建 Pull Request
![implement_issue](./docs/images/implement_issue.jpg)
3. 评论&关闭Issue
![comment_and_close_issue](./docs/images/comment_and_close_issue.jpg)
</details>

## 安装

### 前提条件

- Go 1.23.0 或更高版本
- 拥有访问令牌的 Gitee 账户，[前往获取](https://gitee.com/profile/personal_access_tokens)

### 从源代码构建

1. 克隆仓库：
   ```bash
   git clone https://gitee.com/oschina/mcp-gitee.git
   cd mcp-gitee
   ```

2. 构建项目：
   ```bash
   make build
   ```
   将 ./bin/mcp-gitee 移动至系统环境变量
   
### 使用 go install 安装
   ```bash
   go install gitee.com/oschina/mcp-gitee@latest
   ```

## 使用方法

检查 mcp-gitee 版本：

```bash
mcp-gitee --version
```

### MCP Hosts 配置

例如，以 Windsurf、Cursor 为例，Claude Desktop、Cline、RooCode 都是类似的。

**Cursor**:

stdio mode
```bash
mcp-gitee -token <Your Personal Access Token>
```

sse mode
```bash
mcp-gitee -transport sse -token <Your Personal Access Token>
```

**Windsurf**:
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

### 命令行选项

- `-token`：Gitee 访问令牌
- `-api-base`：Gitee API 基础 URL（默认：https://gitee.com/api/v5）
- `-version`：显示版本信息
- `-transport`：传输类型（stdio 或 sse，默认：stdio）
- `-sse-address`：启动 SSE 服务器的主机和端口（默认：localhost:8000）

### 环境变量

您也可以使用环境变量配置服务器：

- `GITEE_ACCESS_TOKEN`：Gitee 访问令牌
- `GITEE_API_BASE`：Gitee API 基础 URL

## 许可证

本项目采用 MIT 许可证。有关更多详细信息，请参阅 [LICENSE](LICENSE) 文件。

## 可用工具

服务器提供了各种与 Gitee 交互的工具：

| 工具                          | 类别 | 描述               |
|-----------------------------|------|------------------|
| **list_user_repos**         | 仓库 | 列出用户授权的仓库        |
| **get_file_content**        | 仓库 | 获取仓库中文件的内容       |
| **create_user_repo**        | 仓库 | 创建用户仓库           |
| **create_org_repo**         | 仓库 | 创建组织仓库           |
| **create_enter_repo**       | 仓库 | 创建企业仓库           |
| **create_release**          | 仓库 | 为仓库创建发行版         |
| **list_releases**           | 仓库 | 列出仓库发行版          |
| **list_repo_pulls**         | Pull Request | 列出仓库中的拉取请求       |
| **merge_pull**              | Pull Request | 合并拉取请求           |
| **create_pull**             | Pull Request | 创建拉取请求           |
| **update_pull**             | Pull Request | 更新拉取请求           |
| **get_pull_detail**         | Pull Request | 获取拉取请求的详细信息      |
| **comment_pull**            | Pull Request | 评论拉取请求           |
| **list_pull_comments**      | Pull Request | 列出拉取请求的所有评论      |
| **create_issue**            | Issue | 创建 Issue         |
| **update_issue**            | Issue | 更新 Issue         |
| **get_repo_issue_detail**   | Issue | 获取仓库 Issue 的详细信息 |
| **list_repo_issues**        | Issue | 列出仓库 Issue       |
| **comment_issue**           | Issue | 评论 Issue         |
| **list_issue_comments**     | Issue | 列出 Issue 的评论     |
| **get_user_info**           | 用户 | 获取当前认证用户信息 |
| **list_user_notifications** | 通知 | 列出用户通知           |

## 贡献

我们欢迎开源社区的贡献！如果您想为这个项目做出贡献，请按照以下指南操作：

1. Fork 这个仓库。
2. 为您的功能或 bug 修复创建一个新分支。
3. 进行更改，并确保代码有良好的文档。
4. 提交一个 pull request，并附上清晰的更改描述。

更多信息，请参阅 [CONTRIBUTING](CONTRIBUTING.md) 文件。
