package service

func PrintIssues(broker *Broker, writer *Writer, projectName string, sprintName string) error {
	issues, err := broker.GetIssues(
		projectName,
		sprintName,
		GetIssuesQuery{},
	)
	if err != nil {
		return err
	}
	err = writer.PrintStories(issues.Issues...)
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
	err = writer.PrintStories(issues.Issues...)
	if err != nil {
		return err
	}
	return nil
}
