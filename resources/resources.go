package resources

import (
	"errors"
	"net/http"

	"github.com/gocraft/dbr"
)

var (
	ErrDatabaseQuery = errors.New("database query error")
	ErrNotExist      = errors.New("resource not found")
)

type Resource interface {
	GetDbh() *dbr.Connection
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
