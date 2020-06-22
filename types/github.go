package types

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request error: %s", res.Status)
	}

	var result *IssueSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func StartSearch() {
	// Build sample: main repo:Ythosa/where-is is:open
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	//SimplePrint(result)
	//TemplateTextPrint(result)
	TemplateHTMLPrint(os.Stdout, result)
}

func ServerSearch(out io.Writer, args []string) {
	result, err := SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}
	TemplateHTMLPrint(out, result)
}

func daysAgo(date time.Time) int {
	return int(time.Since(date).Hours() / 24)
}

func TemplateTextPrint(result *IssueSearchResult) {
	const templ = `Total {{.TotalCount}} issues:
	{{range .Items}}
-------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}`

	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func TemplateHTMLPrint(out io.Writer, result *IssueSearchResult) {
	var issueList = template.Must(template.New("issuelist").Parse(
		`<h1>Total: {{.TotalCount}} issues</h1>
<table>
	<tr style=’text-align: left’>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr> 
	{{range .Items}}
	<tr>
		<td><a href= '{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr> 
	{{end}}
</table>`))

	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}

func SimplePrint(result *IssueSearchResult) {
	fmt.Printf("Total: %d issues", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
