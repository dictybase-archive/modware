package stock

import (
	"database/sql"
	"net/http"
)

type Resource struct {
	Dbh *sql.DB
}

func (stock *Resource) GetOrder(w http.ResponseWriter, r *http.Request) {
}

func (stock *Resource) GetAllOrders(w http.ResponseWriter, r *http.Request) {
}

func (stock *Resource) CreateOrder(w http.ResponseWriter, r *http.Request) {
}

func (stock *Resource) UpdateOrder(w http.ResponseWriter, r *http.Request) {
}

func (stock *Resource) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}
