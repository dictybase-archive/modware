// Package modwaretest provides common constants and functions for unit testing
package modwaretest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Jeffail/gabs"
	"github.com/dictyBase/go-middlewares/middlewares/router"
	"github.com/dictyBase/modware/resources"
	"github.com/julienschmidt/httprouter"
	"github.com/lann/builder"
)

const (
	// APIHost is the default http host for tesing
	APIHost = "https://api.dictybase.org"
	// PathPrefix is the default prefix for appending to the API host
	PathPrefix = "1.0"
	// PubID is publication id for testing
	PubID = "99"
)

// IndentJSON uniformly indent the json byte
func IndentJSON(b []byte) []byte {
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", " ")
	return bytes.TrimSpace(out.Bytes())
}

// APIServer returns a server URL
func APIServer() string {
	return fmt.Sprintf("%s/%s", APIHost, PathPrefix)
}

// MatchJSON compares actual and expected json
func MatchJSON(actual []byte, data interface{}) error {
	expected, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if bytes.Compare(IndentJSON(actual), IndentJSON(expected)) != 0 {
		return fmt.Errorf("actual %s and expected json %s are different", string(IndentJSON(actual)), string(IndentJSON(expected)))
	}
	return nil
}

var (
	// HTTPExpectBuilder provides methods for incremental building of http configuration
	// to construct a Request object
	HTTPExpectBuilder = builder.Register(httpExpectBuilder{}, httpExpect{}).(httpExpectBuilder)
	// HTTPRequestBuilder provides methods for incremental building of Request object to receive
	// a Response object
	HTTPRequestBuilder = builder.Register(httpRequestBuilder{}, httpRequest{}).(httpRequestBuilder)
	// HTTPResponseBuilder provides methods for incremental testing of http response and json object.
	HTTPResponseBuilder = builder.Register(httpResponseBuilder{}, httpResponse{}).(httpResponseBuilder)
)

// Reporter interface is used for reporting test failures
type Reporter interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Log(...interface{})
	Logf(string, ...interface{})
}

type httpExpect struct {
	host     string
	resource resources.Resource
	reporter Reporter
}

type httpExpectBuilder builder.Builder

// NewHTTP sets up the Reporter and host
func (b httpExpectBuilder) NewHTTP(rep Reporter, host string) httpExpectBuilder {
	builder.Set(b, "reporter", rep)
	return builder.Set(b, "host", host).(httpExpectBuilder)
}

// Host sets the http host
func (b httpExpectBuilder) Host(host string) httpExpectBuilder {
	return builder.Set(b, "host", host).(httpExpectBuilder)
}

// Resource sets the resources.Resource implementation
func (b httpExpectBuilder) Resource(r resources.Resource) httpExpectBuilder {
	return builder.Set(b, "resource", r).(httpExpectBuilder)
}

// Get configures Request to execute a http GET request
func (b httpExpectBuilder) Get(path string) httpRequestBuilder {
	e := builder.GetStruct(b).(httpExpect)
	r := httptest.NewRequest("GET", fmt.Sprintf("%s/%s", e.host, path), nil)
	w := httptest.NewRecorder()
	return HTTPRequestBuilder.reporter(e.reporter).request(r).response(w).handlerFunc(e.resource.Get)
}

type httpRequest struct {
	req       *http.Request
	res       *httptest.ResponseRecorder
	params    []httprouter.Params
	handlerFn http.HandlerFunc
	reporter  Reporter
}

type httpRequestBuilder builder.Builder

func (b httpRequestBuilder) reporter(rep Reporter) httpRequestBuilder {
	return builder.Set(b, "reporter", rep).(httpRequestBuilder)
}

func (b httpRequestBuilder) handlerFunc(fn http.HandlerFunc) httpRequestBuilder {
	return builder.Set(b, "handlerFn", fn).(httpRequestBuilder)
}

func (b httpRequestBuilder) request(r *http.Request) httpRequestBuilder {
	return builder.Set(b, "req", r).(httpRequestBuilder)
}

func (b httpRequestBuilder) response(w *httptest.ResponseRecorder) httpRequestBuilder {
	return builder.Set(b, "res", w).(httpRequestBuilder)
}

// AddRouterParam add key and value to httprouter's parameters
func (b httpRequestBuilder) AddRouterParam(key, value string) httpRequestBuilder {
	_, ok := builder.Get(b, "params")
	if !ok {
		var p httprouter.Params
		p = append(p, httprouter.Param{Key: key, Value: value})
		return builder.Set(b, "params", p).(httpRequestBuilder)
	}
	return builder.Append(b, "params", httprouter.Param{Key: key, Value: value}).(httpRequestBuilder)
}

// Expect gets the Response object for further testing
func (b httpRequestBuilder) Expect() httpResponseBuilder {
	httpr := builder.GetStruct(b).(httpRequest)
	ctx := context.WithValue(context.Background(), router.ContextKeyParams, httpr.params)
	httpr.handlerFn(httpr.res, httpr.req.WithContext(ctx))
	return HTTPResponseBuilder.reporter(httpr.reporter).response(httpr.res)
}

type httpResponse struct {
	reporter Reporter
	res      httptest.ResponseRecorder
	failed   bool
}

type httpResponseBuilder builder.Builder

func (b httpResponseBuilder) reporter(rep Reporter) httpResponseBuilder {
	return builder.Set(b, "reporter", rep).(httpResponseBuilder)
}

func (b httpResponseBuilder) response(w *httptest.ResponseRecorder) httpResponseBuilder {
	return builder.Set(b, "res", w).(httpResponseBuilder)
}

// Status matches the expected and actual http status
func (b httpResponseBuilder) Status(status int) httpResponseBuilder {
	failure := false
	v, _ := builder.Get(b, "res")
	w := v.(httptest.ResponseRecorder)
	if w.Code != status {
		failure = true
		rv, _ := builder.Get(b, "reporter")
		rep := rv.(Reporter)
		rep.Errorf("actual http status %d did not match with expected status %d\n", status, w.Code)
	}
	return builder.Set(b, "failed", failure).(httpResponseBuilder)
}

// JSON return a container type for introspecting json response
func (b httpResponseBuilder) JSON() *gabs.Container {
	v, _ := builder.Get(b, "res")
	w := v.(httptest.ResponseRecorder)
	cont, err := gabs.ParseJSON(w.Body.Bytes())
	if err != nil {
		rv, _ := builder.Get(b, "reporter")
		rep := rv.(Reporter)
		rep.Fatalf("unable to parse json from response body %s\n", err)
	}
	return cont
}
