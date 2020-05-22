package service

import (
	"errors"
	"strconv"
)

func PrintIssues(broker *Broker, writer *Writer, projectName string, sprintName string) error {
	issues, err := broker.GetIssues(
		projectName,
		sprintName,
		GetIssuesQuery{},
	)
	if err != nil {
		return err
	}
	return writer.PrintIssues(issues.Issues...)
}

func PrintIssue(broker *Broker, writer *Writer, projectName string, sprintName string, issueId int) error {
	issues, err := broker.GetIssues(
		projectName,
		sprintName,
		GetIssuesQuery{IssueId: &issueId},
	)
	if err != nil {
		return err
	}
	if len(issues.Issues) == 0 {
		return errors.New("not found issue from issue id " + strconv.Itoa(issueId))
	}
	return writer.PrintIssue(issues.Issues[0])
}
