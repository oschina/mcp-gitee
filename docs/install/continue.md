Continue hub 助手使用 config.yaml 规范定义。本地助手也可以通过放置在全局 .continue 文件夹中的 YAML 文件 config.yaml 进行配置（Mac 上为 ~/.continue，Windows 上为 %USERPROFILE%\.continue）

配置示例：

```yml
mcpServers:
  - name: gitee
    command: mcp-gitee
    args:
      - --token
      - <your personal token>
```