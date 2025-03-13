package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gitee.com/oschina/mcp-gitee/operations/issues"
	"gitee.com/oschina/mcp-gitee/operations/notifications"
	"gitee.com/oschina/mcp-gitee/operations/pulls"
	"gitee.com/oschina/mcp-gitee/operations/repository"
	"gitee.com/oschina/mcp-gitee/operations/users"
	"gitee.com/oschina/mcp-gitee/utils"

	"github.com/mark3labs/mcp-go/server"
)

var (
	// Version gitee mcp server version
	Version = "0.1.1"
)

func newMCPServer() *server.MCPServer {
	return server.NewMCPServer(
		"mcp-gitee",
		Version,
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
}

func addTools(s *server.MCPServer) {
	// Repository Tools
	s.AddTool(repository.ListUserReposTool, repository.ListUserReposHandler)
	s.AddTool(repository.GetFileContentTool, repository.GetFileContentHandler)
	s.AddTool(repository.NewCreateRepoTool(repository.CreateUserRepo), repository.CreateRepoHandleFunc(repository.CreateUserRepo))
	s.AddTool(repository.NewCreateRepoTool(repository.CreateOrgRepo), repository.CreateRepoHandleFunc(repository.CreateOrgRepo))
	s.AddTool(repository.NewCreateRepoTool(repository.CreateEnterRepo), repository.CreateRepoHandleFunc(repository.CreateEnterRepo))
	s.AddTool(repository.CreateReleaseTool, repository.CreateReleaseHandleFunc)
	s.AddTool(repository.ListReleasesTool, repository.ListReleasesHandleFunc)

	// Pulls Tools
	s.AddTool(pulls.NewListPullsTool(pulls.ListRepoPullsToolName), pulls.ListPullsHandleFunc(pulls.ListRepoPullsToolName))
	s.AddTool(pulls.CreatePullTool, pulls.CreatePullHandleFunc)
	s.AddTool(pulls.UpdatePullTool, pulls.UpdatePullHandleFunc)
	s.AddTool(pulls.GetPullDetailTool, pulls.GetPullDetailHandleFunc)
	s.AddTool(pulls.CommentPullTool, pulls.CommentPullHandleFunc)
	s.AddTool(pulls.MergePullTool, pulls.MergePullHandleFunc)
	s.AddTool(pulls.ListPullCommentsTool, pulls.ListPullCommentsHandleFunc)

	// Issues Tools
	s.AddTool(issues.CreateIssueTool, issues.CreateIssueHandleFunc)
	s.AddTool(issues.UpdateIssueTool, issues.UpdateIssueHandleFunc)
	s.AddTool(issues.NewGetIssueDetailTool(issues.GetRepoIssueDetailToolName), issues.GetIssueDetailHandleFunc(issues.GetRepoIssueDetailToolName))
	s.AddTool(issues.NewListIssuesTool(issues.ListRepoIssuesToolName), issues.ListIssuesHandleFunc(issues.ListRepoIssuesToolName))
	s.AddTool(issues.CommentIssueTool, issues.CommentIssueHandleFunc)
	s.AddTool(issues.ListIssueCommentsTool, issues.ListIssueCommentsHandleFunc)

	// Notifications Tools
	s.AddTool(notifications.ListUserNotificationsTool, notifications.ListUserNotificationsHandler)

	// Users Tools
	s.AddTool(users.GetUserInfoTool, users.GetUserInfoHandleFunc())
}

func run(transport, addr string) error {
	s := newMCPServer()
	addTools(s)

	switch transport {
	case "stdio":
		if err := server.ServeStdio(s); err != nil {
			return err
		}
	case "sse":
		srv := server.NewSSEServer(s, "http://"+addr)
		log.Printf("SSE server listening on %s", addr)
		if err := srv.Start(addr); err != nil {
			return fmt.Errorf("server error: %v", err)
		}
	default:
		return fmt.Errorf(
			"invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}
	return nil
}

func main() {
	accessToken := flag.String("token", "", "Gitee access token")
	apiBase := flag.String("api-base", "", "Gitee API base URL (default: https://gitee.com/api/v5)")
	showVersion := flag.Bool("version", false, "Show version information")
	transport := flag.String("transport", "stdio", "Transport type (stdio or sse)")
	addr := flag.String("sse-address", "localhost:8000", "The host and port to start the sse server on")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Gitee MCP Server\n")
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	if *accessToken != "" {
		utils.SetGiteeAccessToken(*accessToken)
	}

	if *apiBase != "" {
		utils.SetApiBase(*apiBase)
	}

	if err := run(*transport, *addr); err != nil {
		panic(err)
	}

}
