package publication

import (
	"database/sql"
	"net/http"
)

type Publication struct {
	Dbh *sql.DB
}

func (pub *Publication) GetDbh() *sql.DB {
	return pub.Dbh
}

func (pub *Publication) Get(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Create(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Update(w http.ResponseWriter, r *http.Request) {

}

func (pub *Publication) Delete(w http.ResponseWriter, r *http.Request) {

}

type Author struct {
	Dbh *sql.DB
}

func (a *Author) GetDbh() *sql.DB {
	return a.Dbh
}

func (a *Author) Get(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) Create(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) Update(w http.ResponseWriter, r *http.Request) {

}

func (a *Author) Delete(w http.ResponseWriter, r *http.Request) {

}
