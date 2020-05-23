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

func (w *Writer) PrintIssues(issue ...Issue) error {
	rows := make([][]string, len(issue))
	for _, v := range issue {
		rows = append(rows, []string{strconv.Itoa(v.ID), v.Subject, v.AssignedTo.Name})
	}

	table := w.table()
	table.SetHeader([]string{"ID", "SUBJECT", "ASSIGNED"})
	table.AppendBulk(rows)
	table.Render()
	return nil
}

func (w *Writer) PrintIssue(issue Issue) error {
	w.write("[INFO]\n")
	infoSection := uitable.New()
	infoSection.AddRow("ID: ", issue.ID)
	infoSection.AddRow("SUBJECT: ", issue.Subject)
	infoSection.AddRow("ASSIGNED: ", issue.AssignedTo.Name)
	w.write(infoSection.String())
	w.write("\n\n[DESCRIPTION]\n")
	w.write(replaceLineFeedCode(issue.Description))
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
