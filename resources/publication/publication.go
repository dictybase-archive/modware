package publication

import (
	"database/sql"
	"net/http"
)

type Resource struct {
	Dbh *sql.DB
}

func (pub *Resource) GetPublication(w http.ResponseWriter, r *http.Response) {

}

func (pub *Resource) GetAllPublictions(w http.ResponseWriter, r *http.Response) {

}

func (pub *Resource) CreatePublication(w http.ResponseWriter, r *http.Response) {

}

func (pub *Resource) UpdatePublication(w http.ResponseWriter, r *http.Response) {

}

func (pub *Resource) DeletePublication(w http.ResponseWriter, r *http.Response) {

}
