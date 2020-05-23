package service

import (
	"github.com/gosuri/uitable"
	"github.com/olekukonko/tablewriter"
	"io"
	"strconv"
	"strings"
)

type Writer struct {
	Out io.Writer
}

func (w *Writer) PrintIssues(issues []Issue) error {
	rows := make([][]string, len(issues))
	for _, v := range issues {
		rows = append(rows, []string{strconv.Itoa(v.ID), v.Status.Name, v.Subject, v.AssignedTo.Name})
	}

	table := w.table()
	table.SetHeader([]string{"ID", "STATUS", "SUBJECT", "ASSIGNED"})
	table.AppendBulk(rows)
	table.Render()
	return nil
}

func (w *Writer) PrintIssue(issue Issue, children []Issue) error {
	w.write("[INFO]\n")
	infoSection := uitable.New()
	infoSection.AddRow("ID: ", issue.ID)
	infoSection.AddRow("STATUS: ", issue.Status.Name)
	infoSection.AddRow("SUBJECT: ", issue.Subject)
	infoSection.AddRow("ASSIGNED: ", issue.AssignedTo.Name)
	w.write(infoSection.String())

	w.write("\n")
	w.write("\n[DESCRIPTION]")
	if desc := replaceLineFeedCode(issue.Description); desc != "" {
		w.write("\n" + desc)
	}

	w.write("\n")
	w.write("\n[CHILDREN]\n")
	if len(issue.Children) != 0 {
		childrenSection := w.table()
		childrenSection.SetHeader([]string{"ID", "STATUS", "SUBJECT", "ASSIGNED"})
		for _, v := range children {
			childrenSection.Append([]string{strconv.Itoa(v.ID), v.Status.Name, v.Subject, v.AssignedTo.Name})
		}
		childrenSection.Render()
	}

	return nil
}

func (w *Writer) table() *tablewriter.Table {
	table := tablewriter.NewWriter(w.Out)
	table.SetBorders(tablewriter.Border{Left: true, Right: true, Top: false, Bottom: false})
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	return table
}

//noinspection GoUnhandledErrorResult
func (w *Writer) write(str string) {
	w.Out.Write([]byte(str))
}

func replaceLineFeedCode(str string) string {
	return strings.NewReplacer("\r\n", "\n").Replace(str)
}
