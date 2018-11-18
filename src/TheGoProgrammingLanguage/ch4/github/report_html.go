package main

import (
	"github"
	"html/template"
	"time"
	"log"
	"os"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
    <tr style='tesxt-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
    <th>Title</th>
    </tr>
{{range .Items}}
    <tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
    </tr>
{{end}}
</table>`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}


func noMust() {
	report, err:= template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)

	if err != nil {
		log.Fatal(err)
	}

	result, err := github.SearchIssues([]string{"12345"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func main() {
	noMust()
}