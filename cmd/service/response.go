// Redmine REST API response type
// autogenerated by https://mholt.github.io/json-to-go/

package service

import (
	"errors"
	"fmt"
	"time"
)

type IssuesResponse struct {
	Issues     []Issue `json:"issues"`
	TotalCount int     `json:"total_count"`
	Offset     int     `json:"offset"`
	Limit      int     `json:"limit"`
}

type IssueResponse struct {
	Issue Issue `json:"issue"`
}

type Issue struct {
	ID      int `json:"id"`
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"project"`
	Tracker struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tracker"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	Priority struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"priority"`
	Author struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"author"`
	AssignedTo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"assigned_to"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	FixedVersion struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"fixed_version"`
	Subject        string    `json:"subject"`
	Description    string    `json:"description"`
	StartDate      string    `json:"start_date"`
	DoneRatio      int       `json:"done_ratio"`
	IsPrivate      bool      `json:"is_private"`
	EstimatedHours float64   `json:"estimated_hours"`
	CreatedOn      time.Time `json:"created_on"`
	UpdatedOn      time.Time `json:"updated_on"`
	ClosedOn       time.Time `json:"closed_on"`
	Children       []struct {
		ID      int `json:"id"`
		Tracker struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tracker"`
		Subject string `json:"subject"`
	} `json:"children"`
}

type ProjectsResponse struct {
	Projects []struct {
		ID             int       `json:"id"`
		Name           string    `json:"name"`
		Identifier     string    `json:"identifier"`
		Description    string    `json:"description"`
		Status         int       `json:"status"`
		IsPublic       bool      `json:"is_public"`
		InheritMembers bool      `json:"inherit_members"`
		CreatedOn      time.Time `json:"created_on"`
		UpdatedOn      time.Time `json:"updated_on"`
	} `json:"projects"`
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

func (r ProjectsResponse) findProjectId(projectName string) (int, error) {
	for _, v := range r.Projects {
		if v.Name == projectName {
			return v.ID, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Not found project id by project name [%s]", projectName))
}

type VersionsResponse struct {
	Versions []struct {
		ID      int `json:"id"`
		Project struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"project"`
		Name          string      `json:"name"`
		Description   string      `json:"description"`
		Status        string      `json:"status"`
		DueDate       interface{} `json:"due_date"`
		Sharing       string      `json:"sharing"`
		WikiPageTitle string      `json:"wiki_page_title"`
		CreatedOn     time.Time   `json:"created_on"`
		UpdatedOn     time.Time   `json:"updated_on"`
	} `json:"versions"`
	TotalCount int `json:"total_count"`
}

func (r VersionsResponse) findVersionId(versionName string) (int, error) {
	for _, v := range r.Versions {
		if v.Name == versionName {
			return v.ID, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Not found version id by version name [%s]", versionName))
}
