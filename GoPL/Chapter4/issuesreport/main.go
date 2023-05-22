package main

import (
	"Chapter4/github/github"
	"log"
	"os"
	"text/template"
	"time"
)

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
func main() {
	q := os.Args[1:]
	result, err := github.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
