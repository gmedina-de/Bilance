package repository

import (
	"homecloud/core/database"
	"homecloud/core/model"
	"reflect"
)

type generic[T model.Model] struct {
	database database.Database
	model    T
}

func Generic[T model.Model](database database.Database, model T) Repository[T] {
	database.AutoMigrate(model)
	return &generic[T]{database: database, model: model}
}

func (r *generic[T]) All() []T {
	var result []T
	r.database.Find(&result)
	return result
}

func (r *generic[T]) Count() int64 {
	var result = new(int64)
	r.database.Model(r.model).Count(result)
	return *result
}

func (r *generic[T]) Limit(limit int, offset int) []T {
	var result []T
	r.database.Model(r.model).Limit(limit).Offset(offset).Find(&result)
	return result
}

func (r *generic[T]) Find(id int64) *T {
	var result T
	r.database.First(&result, id)
	return &result
}

func (r *generic[T]) List(query string, args ...string) []T {
	var result []T
	r.database.Where(query, args).Find(&result)
	return result
}

func (r *generic[T]) Map(query string, args ...string) map[int64]*T {
	var result = make(map[int64]*T)
	list := r.List(query, args...)
	for _, elem := range list {
		result[reflect.ValueOf(elem).FieldByName("Id").Interface().(int64)] = &elem
	}
	return result
}

func (r *generic[T]) Insert(entity any) {
	r.database.Create(entity)
}

func (r *generic[T]) Update(entity any) {
	r.database.Save(entity)
}

func (r *generic[T]) Delete(entity any) {
	r.database.Delete(entity)
}
