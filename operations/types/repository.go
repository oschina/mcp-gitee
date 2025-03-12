package types

type Project struct {
	Id              int             `json:"id"`
	FullName        string          `json:"full_name"`
	HumanName       string          `json:"human_name"`
	Url             string          `json:"url"`
	Namespace       Namespace       `json:"namespace"`
	Owner           BasicUser       `json:"owner"`
	Assigner        BasicUser       `json:"assigner"`
	Description     string          `json:"description"`
	Private         bool            `json:"private"`
	Public          bool            `json:"public"`
	Internal        bool            `json:"internal"`
	Fork            bool            `json:"fork"`
	SshUrl          string          `json:"ssh_url"`
	Recommend       bool            `json:"recommend"`
	Gvp             bool            `json:"gvp"`
	StargazersCount int             `json:"stargazers_count"`
	ForksCount      int             `json:"forks_count"`
	WatchersCount   int             `json:"watchers_count"`
	DefaultBranch   string          `json:"default_branch"`
	Members         []string        `json:"members"`
	PushedAt        string          `json:"pushed_at"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
	Permission      map[string]bool `json:"permission"`
	Status          string          `json:"status"`
	Enterprise      BasicEnterprise `json:"enterprise"`
}

type Namespace struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	HtmlUrl string `json:"html_url"`
}
