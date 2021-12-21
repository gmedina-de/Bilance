package database

import (
	"database/sql"
)

type Database interface {
	Select(table string, resultSlice interface{}, columns string, queryFunc QueryFunc, conditions ...string)
	Insert(table string, model interface{}) sql.Result
	Update(table string, model interface{}) sql.Result
	Delete(table string, model interface{}) sql.Result
	MultipleDelete(table string, conditions ...string) sql.Result
}

type QueryFunc func(row *sql.Rows) any
