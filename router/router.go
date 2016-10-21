package router

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/justinas/alice"
)

const params = "params"

// Wrapper for httprouter
type RouterWrapper struct {
	Router *httprouter.Router
}

// Creates a new instance
func NewRouter() *RouterWrapper {
	return &RouterWrapper{Router: httprouter.New()}
}

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (r *RouterWrapper) Delete(path string, fn http.HandlerFunc) {
	r.Router.DELETE(path, HandlerFunc(fn))
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (r *RouterWrapper) Get(path string, fn http.HandlerFunc) {
	r.Router.GET(path, HandlerFunc(fn))
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (r *RouterWrapper) Head(path string, fn http.HandlerFunc) {
	r.Router.HEAD(path, HandlerFunc(fn))
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (r *RouterWrapper) Options(path string, fn http.HandlerFunc) {
	r.Router.OPTIONS(path, HandlerFunc(fn))
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (r *RouterWrapper) Patch(path string, fn http.HandlerFunc) {
	r.Router.PATCH(path, HandlerFunc(fn))
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (r *RouterWrapper) Post(path string, fn http.HandlerFunc) {
	r.Router.POST(path, HandlerFunc(fn))
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (r *RouterWrapper) Put(path string, fn http.HandlerFunc) {
	r.Router.PUT(path, HandlerFunc(fn))
}

// HandlerFunc accepts the name of a function so you don't have to wrap it with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.HandlerFunc(controller.Index))
// Source: http://nicolasmerouze.com/guide-routers-golang/
func HandlerFunc(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(r.Context(), params, p)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Handler accepts a handler to make it compatible with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.Handler(http.HandlerFunc(controller.Index)))
// Source: http://nicolasmerouze.com/guide-routers-golang/
func Handler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(r.Context(), params, p)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Chain returns handle with chaining using Alice
// Example
//   r.GET("/", router.Chain(resources.GetDataHandler, mw.Logging, mw.Authentication))
func Chain(fn http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	return Handler(alice.New(c...).ThenFunc(fn))
}

// Params returns the Params structure of httprouter
func Params(r *http.Request) httprouter.Params {
	return r.Context().Value(params).(httprouter.Params)
}
