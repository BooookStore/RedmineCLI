package service

import (
	"github.com/gosuri/uitable"
	"io"
	"strings"
)

type Writer struct {
	Out io.Writer
}

func (w *Writer) PrintIssues(issue ...Issue) error {
	table := uitable.New()
	table.MaxColWidth = 500
	table.AddRow("ID", "SUBJECT", "STATUS", "ASSIGNED")
	for _, v := range issue {
		table.AddRow(v.ID, v.Subject, v.Status.Name, v.AssignedTo.Name)
	}
	_, err := w.Out.Write([]byte(table.String()))
	if err != nil {
		return err
	}
	return nil
}

func (w *Writer) PrintIssue(issue Issue) error {
	header := uitable.New()
	header.AddRow("ID", "SUBJECT", "STATUS", "ASSIGNED")
	header.AddRow(issue.ID, issue.Subject, issue.Status.Name, issue.AssignedTo.Name)
	_, err := w.Out.Write([]byte(header.String()))
	if err != nil {
		return err
	}
	sb := strings.Builder{}
	sb.WriteString("\n")
	sb.WriteString("DESCRIPTION\n")
	sb.WriteString(replaceLineFeedCode(issue.Description))
	_, err = w.Out.Write([]byte(sb.String()))
	if err != nil {
		return err
	}
	return nil
}

func replaceLineFeedCode(str string) string {
	return strings.NewReplacer("\r\n", "\n").Replace(str)
}
