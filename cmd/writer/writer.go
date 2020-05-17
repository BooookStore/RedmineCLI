package writer

import (
	"github.com/BooookStore/RedmineCLI/cmd/service"
	"github.com/gosuri/uitable"
	"io"
)

type Writer struct {
	Out io.Writer
}

func (w *Writer) PrintStories(issue ...service.Issue) error {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("ID", "SUBJECT")
	for _, v := range issue {
		table.AddRow(v.ID, v.Subject)
	}
	_, err := w.Out.Write([]byte(table.String()))
	if err != nil {
		return err
	}
	return nil
}
