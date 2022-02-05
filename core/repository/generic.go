package repository

import (
	"homecloud/core/database"
	"homecloud/core/model"
	"reflect"
)

type Generic[T model.Model] struct {
	database database.Database
	model    T
}

func NewGeneric[T model.Model](database database.Database, model T) Repository[T] {
	database.AutoMigrate(model)
	return &Generic[T]{database: database, model: model}
}

func (r *Generic[T]) All() []T {
	var result []T
	r.database.Find(&result)
	return result
}

func (r *Generic[T]) Find(id int64) *T {
	var result []T
	r.database.Find(&result, id)
	if len(result) > 0 {
		return &result[0]
	} else {
		return nil
	}
}

func (r *Generic[T]) List(query string, args ...string) []T {
	var result []T
	r.database.Where(query, args).Find(&result)
	return result
}

func (r *Generic[T]) Map(query string, args ...string) map[int64]*T {
	var result = make(map[int64]*T)
	list := r.List(query, args...)
	for _, elem := range list {
		result[reflect.ValueOf(elem).FieldByName("Id").Interface().(int64)] = &elem
	}
	return result
}

func (r *Generic[T]) Insert(entity any) {
	r.database.Create(entity)
}

func (r *Generic[T]) Update(entity any) {
	r.database.Save(entity)
}

func (r *Generic[T]) Delete(entity any) {
	r.database.Delete(entity)
}
