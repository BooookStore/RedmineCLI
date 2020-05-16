package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestBroker_GetIssues(t *testing.T) {
	// setup testClient
	testClient := NewTestClient(func(req *http.Request) *http.Response {
		switch req.URL.String() {
		case "http://redmine.test.com/projects.json":
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body: ioutil.NopCloser(bytes.NewBufferString(`
						{ "projects": 
							[ {
								"id": 1,
								"name": "SampleProject" 
							} ]
						}`)),
			}
		case "http://redmine.test.com/projects/sampleproject/versions.json":
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body: ioutil.NopCloser(bytes.NewBufferString(`
						{
							"versions": [
								{
									"id": 1,
									"name":	"SampleVersion1"
								}
							]
						}`)),
			}
		case "http://redmine.test.com/issues.json?fix_version_id=1&project_id=1":
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{ "total_count": 1 }`)),
			}
		}

		t.Error(errors.New(fmt.Sprintf("Illegal url access [%s]", req.URL.String())))
		return nil
	})

	u, _ := url.Parse("http://redmine.test.com")
	mockRESTClient := &RESTClient{u, "abcdefg", testClient}

	// setup
	b := &Broker{mockRESTClient}

	// execute
	issues, err := b.GetIssues("SampleProject", "SampleVersion1")

	// verify
	assert.Nil(t, err)
	assert.Equal(t, issues.TotalCount, 1)
}
