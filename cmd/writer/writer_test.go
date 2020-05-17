package writer

import (
	"bytes"
	"github.com/BooookStore/RedmineCLI/cmd/service"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestPrintStories(t *testing.T) {
	// setup
	stdout := new(bytes.Buffer)
	writer := Writer{Out: stdout}

	// execute
	err := writer.printStories(
		service.Issue{
			ID:      1,
			Subject: "FirstIssue",
		},
		service.Issue{
			ID:      2,
			Subject: "SecondIssue",
		},
	)

	// verify
	assert.Nil(t, err)
	output, _ := ioutil.ReadAll(stdout)
	assert.Equal(t, `1	FirstIssue
2	SecondIssue
`, string(output))
}
