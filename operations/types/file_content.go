package types

type FileContent struct {
	Type    string `json:"type"`
	Size    int    `json:"size"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	Sha     string `json:"sha"`
	Content string `json:"content"`
}
