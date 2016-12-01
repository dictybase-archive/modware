package resources

import "database/sql"

type Resource interface {
	GetDbh() *sql.DB
	Get()
	GetAll()
	Create()
	Update()
	Delete()
}
