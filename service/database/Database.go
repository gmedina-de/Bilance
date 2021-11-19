package database

import "database/sql"

type Database interface {
	Query(query string) (*sql.Rows, error)
	Insert(model interface{})
	Update(model interface{})
	Delete(table string, id string)
}
