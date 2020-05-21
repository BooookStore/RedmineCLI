package service

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestPrintStories(t *testing.T) {
	// setup
	stdout := new(bytes.Buffer)
	writer := Writer{Out: stdout}

	// execute
	err := writer.PrintStories(
		Issue{
			ID:      1,
			Subject: "FirstIssue",
		},
		Issue{
			ID:      2,
			Subject: "SecondIssue",
		},
	)

	// verify
	assert.Nil(t, err)
	output, _ := ioutil.ReadAll(stdout)
	assert.True(t, strings.Contains(string(output), "1"))
	assert.True(t, strings.Contains(string(output), "FirstIssue"))
	assert.True(t, strings.Contains(string(output), "2"))
	assert.True(t, strings.Contains(string(output), "SecondIssue"))
}
