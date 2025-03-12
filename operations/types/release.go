package types

type Release struct {
	Id              int       `json:"id"`
	Author          BasicUser `json:"author"`
	Name            string    `json:"name"`
	Body            string    `json:"body"`
	Prerelease      bool      `json:"prerelease"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	CreatedAt       string    `json:"created_at"`
	Assets          []Asset   `json:"assets"`
}

type Asset struct {
	BrowserDownloadUrl string `json:"browser_download_url"`
	Name               string `json:"name"`
}
