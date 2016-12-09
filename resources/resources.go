package resources

import (
	"errors"
	"net/http"

	"github.com/gocraft/dbr"
	"github.com/manyminds/api2go/jsonapi"
)

var (
	ErrDatabaseQuery = errors.New("database query error")
	ErrNotExist      = errors.New("resource not found")
)

// Interface for every http resource to implement
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
	// Gets a jsonapi.ServerInformation implementing interface
	GetApiServerInfo() jsonapi.ServerInformation
}

// Type that implements jsonapi.ServerInformation interface
type ApiServer struct {
	BaseUrl string
	Prefix  string
}

func (server *ApiServer) GetBaseURL() string {
	return server.BaseUrl
}

func (server *ApiServer) GetPrefix() string {
	return server.Prefix
}
