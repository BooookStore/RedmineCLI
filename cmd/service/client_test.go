package service

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_Get(t *testing.T) {
	const (
		Url    = "http://redmine.test.com"
		Path   = "/issues.json"
		ApiKey = "api_key_value"
	)

	// setup testClient
	const body = `{ "id": 1, "project": { "id": 1, "name": "SampleProject" } }`
	testClient := NewTestClient(func(req *http.Request) *http.Response {

		// verify accessed full path, api key
		assert.Equal(t, Url+Path, req.URL.String())
		assert.Equal(t, ApiKey, req.Header.Get("X-Redmine-API-Key"))

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
			Header:     make(http.Header),
		}
	})

	// setup client
	client := RESTClient{Url, ApiKey, testClient}

	// execute
	var result interface{}
	err := client.Get(Path, &result)

	// verify
	assert.Nil(t, err)
	actualID := result.(map[string]interface{})["id"].(float64)
	assert.Equal(t, float64(1), actualID)
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
