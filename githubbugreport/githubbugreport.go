package main

import (
	"log"
	"net/http"
	"os"
	"sanwakuwaku/go/issues/githubclient"

	"html/template"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Milestone</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  {{if .Milestone}}
  <td>{{.Milestone.State}}	<a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
  {{ else }}
  <td>-</td>
  {{ end }}
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	result, err := githubclient.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
