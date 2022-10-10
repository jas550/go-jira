package main

import (
	"context"
	jira "github.com/andygrunwald/go-jira/v2/cloud"
	"log"
)

func main() {

	jiraClient, _ := jira.NewClient("https://mattermost.atlassian.net/", nil)

	categories, resp, err := jiraClient.StatusCategory.GetList(context.TODO())
	if err != nil {
		log.Println(resp.StatusCode)
		panic(err)
	}

	for _, statusCategory := range categories {
		log.Println(statusCategory)
	}

	category, resp, err := jiraClient.StatusCategory.Get(context.TODO(), "1")
	if err != nil {
		log.Println(resp.StatusCode)
		panic(err)
	}

	log.Println(category)
}
