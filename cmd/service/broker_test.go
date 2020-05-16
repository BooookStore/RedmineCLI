package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBroker_GetIssues(t *testing.T) {
	// setup
	broker := &Broker{Client: &mockClient{}}

	// execute
	issues, err := broker.GetIssues()
	assert.Nil(t, err)
	assert.Equal(t, 1, issues.TotalCount)
}

type mockClient struct {
}

func (m *mockClient) Get(_ string, result interface{}) error {
	result.(*Issues).TotalCount = 1
	return nil
}
