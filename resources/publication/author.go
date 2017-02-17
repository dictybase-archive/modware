package publication

import (
	"fmt"
	"net/http"

	"github.com/dictyBase/apihelpers/apherror"
	"github.com/gocraft/dbr"
)

type Author struct {
	Dbh        *dbr.Connection
	PathPrefix string
}

func (a *Author) GetDbh() *dbr.Connection {
	return a.Dbh
}

func (a *Author) Get(w http.ResponseWriter, r *http.Request) {
	apherror.JSONAPIError(
		w,
		apherror.ErrNotFound.New(
			fmt.Sprintf("requested resource %s not found", r.URL.Path),
		),
	)
}

func (a *Author) GetAll(w http.ResponseWriter, r *http.Request) {
	apherror.JSONAPIError(
		w,
		apherror.ErrNotFound.New(
			fmt.Sprintf("requested resource %s not found", r.URL.Path),
		),
	)
}

func (a *Author) Create(w http.ResponseWriter, r *http.Request) {
	apherror.JSONAPIError(
		w,
		apherror.ErrNotFound.New(
			fmt.Sprintf("requested resource %s not found", r.URL.Path),
		),
	)
}

func (a *Author) Update(w http.ResponseWriter, r *http.Request) {
	apherror.JSONAPIError(
		w,
		apherror.ErrNotFound.New(
			fmt.Sprintf("requested resource %s not found", r.URL.Path),
		),
	)
}

func (a *Author) Delete(w http.ResponseWriter, r *http.Request) {
	apherror.JSONAPIError(
		w,
		apherror.ErrNotFound.New(
			fmt.Sprintf("requested resource %s not found", r.URL.Path),
		),
	)
}
