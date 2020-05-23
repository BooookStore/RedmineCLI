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
	return writer.PrintIssues(issues.Issues...)
}

func PrintIssue(broker *Broker, writer *Writer, issueId int) error {
	issue, err := broker.GetIssue(issueId)
	if err != nil {
		return err
	}
	return writer.PrintIssue(issue.Issue)
}
