package service

const (
	issuesPath = "issues.json"
)

// Broker is mediate between cli and redmine api
type Broker struct {
	Client Client
}

func (b *Broker) GetIssues() (*IssuesResponse, error) {
	var result IssuesResponse
	err := b.Client.Get(issuesPath, &result)
	return &result, err
}

// Client is redmine client used by Broker
type Client interface {
	Get(string, interface{}) error
}
