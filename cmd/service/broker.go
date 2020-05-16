package service

import (
	"fmt"
	"log"
	"net/url"
)

const (
	issuesPath   = "issues.json"
	projectsPath = "projects.json"
)

// Broker is mediate between cli and redmine api
type Broker struct {
	Client Client
}

func (b *Broker) GetIssues(projectName string) (*IssuesResponse, error) {
	id, err := b.findProjectId(projectName)
	if err != nil {
		return nil, err
	}
	fmt.Println(id)
	var result IssuesResponse
	err = b.Client.Get(issuesPath, &result)
	return &result, err
}

func (b *Broker) findProjectId(projectName string) (int, error) {
	var result ProjectsResponse
	if err := b.Client.Get(projectsPath, &result); err != nil {
		log.Fatal(err)
	}
	return result.findProjectId(projectName)
}

// Client is redmine client used by Broker
type Client interface {
	Get(string, interface{}) error
	Url() url.URL
}
