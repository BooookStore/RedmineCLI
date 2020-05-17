package service

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type RESTClient struct {
	redmineURL string
	apiKey     string
	hclient    *http.Client
}

// TODO remove error return
func NewClient(redmineURL string, apiKey string) (*RESTClient, error) {
	return &RESTClient{redmineURL, apiKey, http.DefaultClient}, nil
}

func (c *RESTClient) Get(path string, result interface{}) error {
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	resp, err := c.hclient.Do(req)
	if err != nil {
		return err
	}
	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTClient) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.redmineURL)
	if err != nil {
		return nil, err
	}
	u, err = u.Parse(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, nil
	}

	// add redmine api key to header
	header := req.Header
	header.Add("X-Redmine-API-Key", c.apiKey)
	req.Header = header

	return req, err
}
