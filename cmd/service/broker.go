package service

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	issuesPath   = "issues.json"
	projectsPath = "projects.json"
	versionsPath = "projects/%s/versions.json"
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
	versionId, err := b.findVersionId(projectName, "SampleVersion1")
	if err != nil {
		return nil, err
	}

	path := func() string {
		values := url.Values{}
		values.Add("project_id", strconv.Itoa(projectId))
		values.Add("fix_version_id", strconv.Itoa(versionId))
		q := url.URL{
			Path:     issuesPath,
			RawQuery: values.Encode(),
		}
		return q.String()
	}

	var result IssuesResponse
	err = b.Client.Get(path(), &result)
	return &result, err
}

func (b *Broker) findProjectId(projectName string) (int, error) {
	var result ProjectsResponse
	if err := b.Client.Get(projectsPath, &result); err != nil {
		return 0, err
	}
	return result.findProjectId(projectName)
}

func (b *Broker) findVersionId(projectName string, versionName string) (int, error) {
	var result VersionsResponse
	if err := b.Client.Get(fmt.Sprintf(versionsPath, strings.ToLower(projectName)), &result); err != nil {
		return 0, err
	}
	return result.findVersionId(versionName)
}

// Client is redmine client used by Broker
type Client interface {
	Get(string, interface{}) error
	Url() url.URL
}
