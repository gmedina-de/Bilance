package repository

import (
	"database/sql"
	"net/http"
)

type Repository interface {
	ModelName() string
	ModelNamePlural() string
	NewEmpty() interface{}
	NewFromRequest(request *http.Request, id int64) interface{}
	NewFromQuery(row *sql.Rows) interface{}
	Find(id int64) interface{}
	List(conditions ...string) interface{}
	Count(conditions ...string) int64
	Insert(entity interface{})
	Update(entity interface{})
	Delete(entity interface{})
}

type idRange []int64

func (r idRange) contains(id int64) bool {
	for _, sliceId := range r {
		if sliceId == id {
			return true
		}
	}
	return false
}

func scanAndPanic(row *sql.Rows, dest ...interface{}) {
	err := row.Scan(dest...)
	if err != nil {
		panic(err)
	}
}
