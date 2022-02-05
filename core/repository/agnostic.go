package repository

import (
	"homecloud/core/database"
	"reflect"
)

type Agnostic struct {
	database database.Database
}

func NewAgnostic(database database.Database) Repository[any] {
	return &Agnostic{database: database}
}

func (r *Agnostic) All() []any {
	var result []any
	r.database.Find(&result)
	return result
}

func (r *Agnostic) Find(id int64) *any {
	var result []any
	r.database.Find(&result, id)
	if len(result) > 0 {
		return &result[0]
	} else {
		return nil
	}
}

func (r *Agnostic) List(query string, args ...string) []any {
	var result []any
	r.database.Where(query, args).Find(&result)
	return result
}

func (r *Agnostic) Map(query string, args ...string) map[int64]*any {
	var result = make(map[int64]*any)
	list := r.List(query, args...)
	for _, elem := range list {
		result[reflect.ValueOf(elem).FieldByName("Id").Interface().(int64)] = &elem
	}
	return result
}

func (r *Agnostic) Insert(entity any) {
	r.database.Create(entity)
}

func (r *Agnostic) Update(entity any) {
	r.database.Save(entity)
}

func (r *Agnostic) Delete(entity any) {
	r.database.Delete(entity)
}
