package service

const (
	issuesPath = "issues.json"
)

// Broker is mediate between cli and redmine api
type Broker struct {
	Client Client
}

func (b *Broker) GetIssues() (*Issues, error) {
	var result Issues
	err := b.Client.Get(issuesPath, &result)
	return &result, err
}

type Client interface {
	Get(string, interface{}) error
}
