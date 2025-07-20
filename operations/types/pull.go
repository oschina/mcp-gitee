package types

type BasicPull struct {
	Id              int        `json:"id"`
	Title           string     `json:"title"`
	Url             string     `json:"html_url"`
	Number          int        `json:"number"`
	State           string     `json:"state"`
	AssigneesNumber int        `json:"assignees_number"`
	TestersNumber   int        `json:"testers_number"`
	Assignees       []Assignee `json:"assignees"`
	Testers         []Assignee `json:"testers"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	ClosedAt        string     `json:"closed_at"`
	MergedAt        string     `json:"merged_at"`
	Creator         BasicUser  `json:"user"`
	Head            Reference  `json:"head"`
	Base            Reference  `json:"base"`
	CanMergeCheck   bool       `json:"can_merge_check"`
	Draft           bool       `json:"draft"`
}

type Reference struct {
	Ref   string `json:"ref"`
	Label string `json:"label"`
	Sha   string `json:"sha"`
}

type Assignee struct {
	BasicUser
	Accept bool `json:"accept"`
}

// PullComment defines the structure for a pull request comment
type PullComment struct {
	Id               int       `json:"id"`
	Body             string    `json:"body"`
	User             BasicUser `json:"user"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
	CommentType      string    `json:"comment_type"`
	CommentId        string    `json:"comment_id"`
	PullUrl          string    `json:"pull_url"`
	OriginalPosition string    `json:"original_position"`
	Path             string    `json:"path"`
	Position         string    `json:"position"`
	CommitId         string    `json:"commit_id"`
	OriginalCommitId string    `json:"original_commit_id"`
	InReplyToId      int       `json:"in_reply_to_id"`
	Url              string    `json:"url"`
}

// DiffFile defines the structure for a pull request diff file
type DiffFile struct {
	Sha       string    `json:"sha"`
	Filename  string    `json:"filename"`
	Status    *string   `json:"status"`
	Additions string    `json:"additions"`
	Deletions string    `json:"deletions"`
	BlobUrl   string    `json:"blob_url"`
	RawUrl    string    `json:"raw_url"`
	Patch     PatchInfo `json:"patch"`
}

// PatchInfo defines the structure for patch information
type PatchInfo struct {
	Diff        string `json:"diff"`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
	TooLarge    bool   `json:"too_large"`
}
