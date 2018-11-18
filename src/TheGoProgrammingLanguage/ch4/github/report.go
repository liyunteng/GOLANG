package main

import (
	"github"
	"text/template"
	"time"
	"log"
	"os"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}} ---------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title|printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`


// var report = template.Must(template.New("issuelist").
// 	Funcs(template.FuncMap{"daysAgo": daysAgo}).
// 	Parse(templ))

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
