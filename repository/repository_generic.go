package repository

import (
	"Bilance/service"
	"database/sql"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Entity interface{}

type GenericRepository[T Entity] interface {
	ModelName() string
	ModelNamePlural() string
	NewEmpty() *T
	NewFromRequest(request *http.Request, id int64) *T
	NewFromQuery(row *sql.Rows) *T
	Find(id int64) *T
	List(conditions ...string) []T
	Count(conditions ...string) int64
	Insert(entity *T)
	Update(entity *T)
	Delete(entity *T)
}

type genericRepository[T Entity] struct {
	database       service.Database
	emptyEntity    T
	newFromQuery   func(row *sql.Rows) *T
	newFromRequest func(request *http.Request, id int64) *T
}

func (r *genericRepository[T]) ModelName() string {
	return strings.ToLower(reflect.TypeOf(r.emptyEntity).Name())
}

func (r *genericRepository[T]) ModelNamePlural() string {
	name := r.ModelName()
	if name[len(name)-1] == 'y' {
		return name[0:len(name)-1] + "ies"
	}
	return name + "s"
}

func (r *genericRepository[T]) NewEmpty() *T {
	emptyEntity := r.emptyEntity
	return &emptyEntity
}

func (r *genericRepository[T]) NewFromQuery(row *sql.Rows) *T {
	return r.newFromQuery(row)
}

func (r *genericRepository[T]) NewFromRequest(request *http.Request, id int64) *T {
	return r.newFromRequest(request, id)
}

func (r *genericRepository[T]) newFromQueryInterface(row *sql.Rows) interface{} {
	return r.newFromQuery(row)
}

func (r *genericRepository[T]) Find(id int64) *T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.newFromQueryInterface, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *genericRepository[T]) List(conditions ...string) []T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.newFromQueryInterface, conditions...)
	return result
}

func (r *genericRepository[T]) Count(conditions ...string) int64 {
	var result []int64
	r.database.Select(r.ModelName(), &result, "COUNT(*)", countQueryFunc, conditions...)
	return result[0]
}

func (r *genericRepository[T]) Insert(entity *T) {
	r.database.Insert(r.ModelName(), entity)
}

func (r *genericRepository[T]) Update(entity *T) {
	r.database.Update(r.ModelName(), entity)
}

func (r *genericRepository[T]) Delete(entity *T) {
	r.database.Delete(r.ModelName(), entity)
}
