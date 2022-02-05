package repository

import (
	"homecloud/core/database"
	"reflect"
)

type agnostic struct {
	database database.Database
}

func Agnostic(database database.Database) Repository[any] {
	return &agnostic{database: database}
}

func (r *agnostic) All() []any {
	var result []any
	r.database.Find(&result)
	return result
}

func (r *agnostic) Find(id int64) *any {
	var result []any
	r.database.Find(&result, id)
	if len(result) > 0 {
		return &result[0]
	} else {
		return nil
	}
}

func (r *agnostic) List(query string, args ...string) []any {
	var result []any
	r.database.Where(query, args).Find(&result)
	return result
}

func (r *agnostic) Map(query string, args ...string) map[int64]*any {
	var result = make(map[int64]*any)
	list := r.List(query, args...)
	for _, elem := range list {
		result[reflect.ValueOf(elem).FieldByName("Id").Interface().(int64)] = &elem
	}
	return result
}

func (r *agnostic) Insert(entity any) {
	r.database.Create(entity)
}

func (r *agnostic) Update(entity any) {
	r.database.Save(entity)
}

func (r *agnostic) Delete(entity any) {
	r.database.Delete(entity)
}
