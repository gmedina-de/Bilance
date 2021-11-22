package repository

import (
	"database/sql"
	"net/http"
)

type Repository interface {
	NewEmpty() interface{}
	NewFromRequest(request *http.Request, id int) interface{}
	NewFromQuery(row *sql.Rows) interface{}
	Find(id string) interface{}
	List(conditions ...string) interface{}
	Insert(interface{})
	Update(interface{})
	Delete(interface{})
}
