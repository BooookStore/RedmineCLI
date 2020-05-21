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
	err = writer.PrintIssues(issues.Issues...)
	if err != nil {
		return err
	}
	return nil
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
	err = writer.PrintIssue(issues.Issues[0])
	if err != nil {
		return err
	}
	return nil
}
