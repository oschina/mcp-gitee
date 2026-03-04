package types

// CommitAuthorInfo is the author/committer info within a commit
type CommitAuthorInfo struct {
	Name  string `json:"name"`
	Date  string `json:"date"`
	Email string `json:"email"`
}

// CommitTree is the tree reference within a commit
type CommitTree struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// CommitDetail is the inner commit object
type CommitDetail struct {
	Author    CommitAuthorInfo `json:"author"`
	Committer CommitAuthorInfo `json:"committer"`
	Message   string           `json:"message"`
	Tree      CommitTree       `json:"tree"`
}

// CommitParent is a parent commit reference
type CommitParent struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// CompareCommit represents a commit entry in a comparison result
type CompareCommit struct {
	Url         string        `json:"url"`
	Sha         string        `json:"sha"`
	HtmlUrl     string        `json:"html_url"`
	CommentsUrl string        `json:"comments_url"`
	Commit      CommitDetail  `json:"commit"`
	Author      *BasicUser    `json:"author"`
	Committer   *BasicUser    `json:"committer"`
	Parents     []CommitParent `json:"parents"`
}

// CompareFile represents a file change in a comparison result
type CompareFile struct {
	Sha        string `json:"sha"`
	Filename   string `json:"filename"`
	Status     string `json:"status"`
	Additions  int    `json:"additions"`
	Deletions  int    `json:"deletions"`
	Changes    int    `json:"changes"`
	BlobUrl    string `json:"blob_url"`
	RawUrl     string `json:"raw_url"`
	ContentUrl string `json:"content_url"`
	Patch      string `json:"patch"`
	Truncated  bool   `json:"truncated"`
}

// CompareResult is the overall response for a branch/tag comparison
type CompareResult struct {
	BaseCommit      CompareCommit   `json:"base_commit"`
	MergeBaseCommit CompareCommit   `json:"merge_base_commit"`
	Commits         []CompareCommit `json:"commits"`
	Files           []CompareFile   `json:"files"`
	Truncated       bool            `json:"truncated"`
}
