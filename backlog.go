package main

import (
	"fmt"
	"log"
	"net/url"

	gobacklog "github.com/griffin-stewie/go-backlog"
)

func issues(space string, token string) (gobacklog.IssueSlice, error) {
	URL, err := url.Parse(fmt.Sprintf("https://%s.backlog.jp", space))
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	client := gobacklog.NewClient(URL, token)

	return client.IssuesWithOption(&gobacklog.IssuesOption{
		Count: 100,
	})
}
