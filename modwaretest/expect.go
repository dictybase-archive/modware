package modwaretest

import (
	"fmt"
	"net/http/httptest"
	"strings"

	"github.com/dictyBase/modware/resources"
)

// ExpectBuilder interface is for incremental building of http configuration
type ExpectBuilder interface {
	Get(string) RequestBuilder
}

// HTTPExpectBuilder implements ExpectBuilder interface
type HTTPExpectBuilder struct {
	reporter Reporter
	host     string
	resource resources.Resource
}

// NewHTTPExpectBuilder is the constructor for HTTPExpectBuilder
func NewHTTPExpectBuilder(rep Reporter, host string, rs resources.Resource) ExpectBuilder {
	return &HTTPExpectBuilder{
		reporter: rep,
		host:     host,
		resource: rs,
	}
}

// Get configures Request to execute a http GET request
func (b *HTTPExpectBuilder) Get(path string) RequestBuilder {
	req := httptest.NewRequest(
		"GET",
		fmt.Sprintf(
			"%s/%s",
			b.host,
			strings.Trim(path, "/"),
		),
		nil,
	)
	return NewHTTPRequestBuilder(b.reporter, req, b.resource.Get)
}
