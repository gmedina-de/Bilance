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

type gRepository[T model.Model[T]] struct {
	database service.Database
	model    T
}

func (r *gRepository[T]) ModelName() string {
	return strings.ToLower(reflect.TypeOf(r.model).Name())
}

func (r *gRepository[T]) ModelNamePlural() string {
	name := r.ModelName()
	if name[len(name)-1] == 'y' {
		return name[0:len(name)-1] + "ies"
	}
	return name + "s"
}

func (r *gRepository[T]) NewEmpty() *T {
	return &r.model
}

func (r *gRepository[T]) NewFromQuery(row *sql.Rows) *T {
	return r.model.FromQuery(row)
}

func (r *gRepository[T]) NewFromRequest(request *http.Request, id int64) *T {
	return r.model.FromRequest(request, id)
}

func (r *gRepository[T]) newFromQueryInterface(row *sql.Rows) interface{} {
	return r.NewFromQuery(row)
}

func (r *gRepository[T]) Find(id int64) *T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.newFromQueryInterface, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *gRepository[T]) List(conditions ...string) []T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.newFromQueryInterface, conditions...)
	return result
}

func (r *gRepository[T]) Count(conditions ...string) int64 {
	var result []int64
	r.database.Select(r.ModelName(), &result, "COUNT(*)", countQueryFunc, conditions...)
	return result[0]
}

func (r *gRepository[T]) Insert(entity *T) {
	r.database.Insert(r.ModelName(), entity)
}

func (r *gRepository[T]) Update(entity *T) {
	r.database.Update(r.ModelName(), entity)
}

func (r *gRepository[T]) Delete(entity *T) {
	r.database.Delete(r.ModelName(), entity)
}
