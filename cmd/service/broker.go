package service

import (
	"log"
	"net/url"
	"strconv"
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
	projectId, err := b.findProjectId(projectName)
	if err != nil {
		return nil, err
	}

	path := func(id int) string {
		values := url.Values{}
		values.Add("project_id", strconv.Itoa(id))
		q := url.URL{}
		q.Path = issuesPath
		q.RawQuery = values.Encode()
		return q.String()
	}

	var result IssuesResponse
	err = b.Client.Get(path(projectId), &result)
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
