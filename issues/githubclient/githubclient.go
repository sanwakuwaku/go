package githubclient

import "time"

const (
	AccessToken     = "token"
	RepoOwner       = "owner"
	RepoName        = "repoName"
	RootEndpoint    = "https://api.github.com"
	SearchIssuesURL = RootEndpoint + "/search/issues"
	IssuesURL       = RootEndpoint + "/repos/" + RepoOwner + "/" + RepoName + "/issues"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Milestone *Milestone `json:"milestone,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	Body      string     // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Number       int
	HTMLURL      string `json:"html_url"`
	State        string
	Title        string
	Desctiption  string
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type IssueRequestParams struct {
	Title     string   `json:"title"`
	Body      string   `json:"body,omitempty"`
	Asignees  []string `json:"assignees,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type EditIssueRequestParams IssueRequestParams
type CreateIssueRequestParams IssueRequestParams
