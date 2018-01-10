package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("backlog-issue", "backlog-issues")

	login = app.Command("login", "login")
	ls    = app.Command("ls", "show issues")
	token = os.Getenv("BACKLOG_TOKEN")
	space = os.Getenv("BACKLOG_SPACE")
)

func main() {
	cmdopt := kingpin.MustParse(app.Parse(os.Args[1:]))
	if token == "" || space == "" {
		log.Fatal("YOU NEED `BACKLOG_TOKEN` and `BACKLOG_SPACE`")
	}

	switch cmdopt {
	case login.FullCommand():
		fmt.Println("login")
	case ls.FullCommand():
		issues, _ := issues(space, token)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"URL", "Summary", "Description"})
		for _, issue := range issues {
			v := []string{
				fmt.Sprintf("https://%s.backlog.jp/view/%s", space, *issue.IssueKey),
				*issue.Summary,
				*issue.Description,
			}
			table.Append(v)
		}
		table.SetBorder(false)
		table.Render()
	}
}
