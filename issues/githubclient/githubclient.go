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

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
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
