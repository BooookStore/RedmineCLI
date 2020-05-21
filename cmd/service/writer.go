package service

import (
	"github.com/gosuri/uitable"
	"io"
)

type Writer struct {
	Out io.Writer
}

func (w *Writer) PrintStories(issue ...Issue) error {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("ID", "SUBJECT", "ASSIGNED")
	for _, v := range issue {
		table.AddRow(v.ID, v.Subject, v.AssignedTo.Name)
	}
	_, err := w.Out.Write([]byte(table.String()))
	if err != nil {
		return err
	}
	return nil
}
