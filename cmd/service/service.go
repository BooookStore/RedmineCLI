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
