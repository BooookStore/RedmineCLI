package writer

import (
	"bufio"
	"fmt"
	"github.com/BooookStore/RedmineCLI/cmd/service"
	"io"
)

type Writer struct {
	Out io.Writer
}

func (w *Writer) printStories(issue ...service.Issue) error {
	bw := bufio.NewWriter(w.Out)
	for _, v := range issue {
		_, err := bw.WriteString(fmt.Sprintf("%v\t%s\n", v.ID, v.Subject))
		if err != nil {
			return err
		}
	}
	err := bw.Flush()
	if err != nil {
		return err
	}
	return nil
}
