package database

import "database/sql"

type Database interface {
	Insert(model interface{})
	Query(query string) (*sql.Rows, error)
}
