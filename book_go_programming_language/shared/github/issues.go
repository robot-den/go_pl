package github

import (
	"fmt"
	"io"
	"time"
)

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

func PrintIssues(w io.Writer, result *IssuesSearchResult) {
	fmt.Printf("%d results found:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("  #%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
