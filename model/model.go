package model

import (
	"database/sql"
	"net/http"
)

type Model[T any] interface {
	Empty() *T
	FromRequest(request *http.Request, id int64) *T
	FromQuery(row *sql.Rows) *T
}

func scanAndPanic(row *sql.Rows, dest ...interface{}) {
	err := row.Scan(dest...)
	if err != nil {
		panic(err)
	}
}
