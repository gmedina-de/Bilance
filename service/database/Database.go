package database

import "database/sql"

type Database interface {
	Select(resultSlice interface{}, queryFunc QueryFunc, conditions ...string)
	Insert(model interface{}) sql.Result
	Update(model interface{}) sql.Result
	Delete(model interface{}) sql.Result
	MultipleDelete(table string, conditions ...string) sql.Result
}

type QueryFunc func(row *sql.Rows) interface{}
