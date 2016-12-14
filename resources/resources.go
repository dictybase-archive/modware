package resources

import (
	"errors"
	"net/http"

	"github.com/dictyBase/go-middlewares/middlewares/pagination"
	"github.com/gocraft/dbr"
	"github.com/manyminds/api2go/jsonapi"
)

var (
	//ErrDatabaseQuery represents database query related errors
	ErrDatabaseQuery = errors.New("database query error")
	//ErrNotExist represents the absence of an HTTP resource
	ErrNotExist = errors.New("resource not found")
	//ErrJSONEncoding represents any json encoding error
	ErrJSONEncoding = errors.New("json encoding error")
	//ErrStructMarshal represents any error with marshalling structure
	ErrStructMarshal = errors.New("structure marshalling error")
)

// Resource is the interface that every http handler have to implement
type Resource interface {
	// Gets the database handler
	GetDbh() *dbr.Connection
	// Handles the http GET for singular resource
	Get(http.ResponseWriter, *http.Request)
	// Handles the http GET for collection resource
	GetAll(http.ResponseWriter, *http.Request)
	// Handles the http POST
	Create(http.ResponseWriter, *http.Request)
	// Handles the http PATCH
	Update(http.ResponseWriter, *http.Request)
	// Handles the http DELETE
	Delete(http.ResponseWriter, *http.Request)
}

// APIServer implements jsonapi.ServerInformation interface
type APIServer struct {
	BaseURL string
	Prefix  string
}

//GetBaseURL returns the base path of the server
func (server *APIServer) GetBaseURL() string {
	return server.BaseURL
}

//GetPrefix returns generic prefix for each server path
func (server *APIServer) GetPrefix() string {
	return server.Prefix
}

// GetAPIServerInfo returns an implementation of jsonapi.ServerInformation
func GetAPIServerInfo(r *http.Request, prefix string) jsonapi.ServerInformation {
	return &APIServer{
		BaseURL: r.URL.Host,
		Prefix:  prefix,
	}
}

//GetPaginationProp returns an instance of pagination.Prop from the request context.
//However, if it's not available, returns one with default value
func GetPaginationProp(r *http.Request) *pagination.Props {
	prop, ok := r.Context().Value(pagination.ContextKeyPagination).(*pagination.Props)
	if ok {
		return prop
	}
	return &pagination.Props{Entries: pagination.DefaultEntries, Current: 1}
}
