package types

type PagedResponse[T any] struct {
	TotalCount int `json:"total_count"`
	Data       []T `json:"data"`
}

type BasicUser struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	HtmlUrl   string `json:"html_url"`
	Remark    string `json:"remark"`
}

type BasicEnterprise struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	HtmlUrl string `json:"html_url"`
}

type BasicProgram struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Assignee    BasicUser `json:"assignee"`
	Author      BasicUser `json:"author"`
}

type EndpointConfig struct {
	UrlTemplate string
	PathParams  []string
}

type IssueTypeDetail struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type BasicLabel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
