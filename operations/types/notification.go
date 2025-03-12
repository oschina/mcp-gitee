package types

type Notification struct {
	Id         int         `json:"id"`
	Content    string      `json:"content"`
	Type       string      `json:"type"`
	Unread     bool        `json:"unread"`
	Mute       bool        `json:"mute"`
	UpdatedAt  string      `json:"updated_at"`
	Url        string      `json:"url"`
	HtmlUrl    string      `json:"html_url"`
	Actor      BasicUser   `json:"actor"`
	Repository Project     `json:"repository"`
	Subject    Subject     `json:"subject"`
	Namespaces []Namespace `json:"namespaces"`
}

type Subject struct {
	Title            string `json:"title"`
	Url              string `json:"url"`
	LatestCommentUrl string `json:"latest_comment_url"`
	Type             string `json:"type"`
}

type NotificationResult struct {
	TotalCount int            `json:"total_count"`
	List       []Notification `json:"list"`
}
