package database

import "database/sql"

type Database interface {
	Query(resultSlice interface{}, queryFunction QueryFunc, conditions ...string)
	Insert(model interface{})
	Update(model interface{})
	Delete(table string, id string)
}

type QueryFunc func(row *sql.Rows) interface{}
