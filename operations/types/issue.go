package types

// BasicIssue 定义了 Issue 的基本结构
type BasicIssue struct {
	Id        int           `json:"id"`
	Number    string        `json:"number"`
	Title     string        `json:"title"`
	User      BasicUser     `json:"user"`
	State     string        `json:"state"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
	Body      string        `json:"body"`
	Labels    []BasicLabel  `json:"labels"`
	Assignee  *BasicUser    `json:"assignee"`
	IssueType string        `json:"issue_type"`
	Program   *BasicProgram `json:"program"`
	HtmlUrl   string        `json:"html_url"`
}

// IssueComment defines the structure for an issue comment
type IssueComment struct {
	Id        int       `json:"id"`
	Body      string    `json:"body"`
	User      BasicUser `json:"user"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}
