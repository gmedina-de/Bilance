package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type GRepository[T model.Model[T]] interface {
	ModelName() string
	ModelNamePlural() string
	NewEmpty() *T
	NewFromQuery(row *sql.Rows) *T
	NewFromRequest(request *http.Request, id int64) *T
	Find(id int64) *T
	List(conditions ...string) []T
	Count(conditions ...string) int64
	Insert(entity *T)
	Update(entity *T)
	Delete(entity *T)
}

type GenericRepository[T model.Model[T]] struct {
	database    service.Database
	constructor T
}

func (r *GenericRepository[T]) ModelName() string {
	return strings.ToLower(reflect.TypeOf(r.constructor).Name())
}

func (r *GenericRepository[T]) ModelNamePlural() string {
	name := r.ModelName()
	if name[len(name)-1] == 'y' {
		return name[0:len(name)-1] + "ies"
	}
	return name + "s"
}

func (r *GenericRepository[T]) NewEmpty() *T {
	return r.constructor.Empty()
}

func (r *GenericRepository[T]) NewFromQuery(row *sql.Rows) *T {
	return r.constructor.FromQuery(row)
}

func (r *GenericRepository[T]) NewFromRequest(request *http.Request, id int64) *T {
	return r.constructor.FromRequest(request, id)
}

func (r *GenericRepository[T]) newFromQueryInterface(row *sql.Rows) interface{} {
	return r.NewFromQuery(row)
}

func (r *GenericRepository[T]) Find(id int64) *T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.newFromQueryInterface, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *GenericRepository[T]) List(conditions ...string) []T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.newFromQueryInterface, conditions...)
	return result
}

func (r *GenericRepository[T]) Count(conditions ...string) int64 {
	var result []int64
	r.database.Select(r.ModelName(), &result, "COUNT(*)", countQueryFunc, conditions...)
	return result[0]
}

func (r *GenericRepository[T]) Insert(entity *T) {
	r.database.Insert(r.ModelName(), entity)
}

func (r *GenericRepository[T]) Update(entity *T) {
	r.database.Update(r.ModelName(), entity)
}

func (r *GenericRepository[T]) Delete(entity *T) {
	r.database.Delete(r.ModelName(), entity)
}
