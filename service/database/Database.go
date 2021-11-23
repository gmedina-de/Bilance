package database

import "database/sql"

type Database interface {
	Select(resultSlice interface{}, queryFunc QueryFunc, conditions ...string)
	Insert(model interface{})
	Update(model interface{})
	Delete(model interface{})
}

type QueryFunc func(row *sql.Rows) interface{}
