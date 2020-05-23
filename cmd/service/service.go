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
	return writer.PrintIssues(issues.Issues)
}

func PrintIssue(broker *Broker, writer *Writer, issueId int) error {
	issue, err := broker.GetIssue(issueId)
	if err != nil {
		return err
	}
	children, err := getChildrenIssue(broker, issue.Issue)
	if err != nil {
		return err
	}
	return writer.PrintIssue(issue.Issue, children)
}

func getChildrenIssue(broker *Broker, issue Issue) ([]Issue, error) {
	var result []Issue
	for _, v := range issue.Children {
		child, err := broker.GetIssue(v.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, child.Issue)
	}
	return result, nil
}
