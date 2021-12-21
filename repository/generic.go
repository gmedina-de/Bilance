package repository

import (
	"Bilance/database"
	"Bilance/model"
	"database/sql"
	"reflect"
	"strconv"
	"strings"
)

type generic[T model.Model] struct {
	database database.Database
	model    T
}

func (r *generic[T]) ModelName() string {
	return strings.ToLower(reflect.TypeOf(r.model).Name())
}

func (r *generic[T]) ModelNamePlural() string {
	name := r.ModelName()
	if name[len(name)-1] == 'y' {
		return name[0:len(name)-1] + "ies"
	}
	return name + "s"
}

func (r *generic[T]) NewEmpty() *T {
	return &r.model
}

func (r *generic[T]) fromQuery(row *sql.Rows) any {
	return r.FromQuery(row)
}

func (r *generic[T]) Find(id int64) *T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.fromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *generic[T]) List(conditions ...string) []T {
	var result []T
	r.database.Select(r.ModelName(), &result, "*", r.fromQuery, conditions...)
	return result
}

func (r *generic[T]) Map(conditions ...string) map[int64]*T {
	var result = make(map[int64]*T)
	list := r.List(conditions...)
	for _, elem := range list {
		result[reflect.ValueOf(elem).FieldByName("Id").Interface().(int64)] = &elem
	}
	return result
}

func (r *generic[T]) Count(conditions ...string) int64 {
	var result []int64
	r.database.Select(r.ModelName(), &result, "COUNT(*)", countQueryFunc, conditions...)
	return result[0]
}

func (r *generic[T]) Insert(entity *T) {
	r.database.Insert(r.ModelName(), entity)
}

func (r *generic[T]) Update(entity *T) {
	r.database.Update(r.ModelName(), entity)
}

func (r *generic[T]) Delete(entity *T) {
	r.database.Delete(r.ModelName(), entity)
}

func countQueryFunc(row *sql.Rows) interface{} {
	var count int64
	row.Scan(row, &count)
	return &count
}
