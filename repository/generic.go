package repository

import (
	"Bilance/database"
	"Bilance/model"
	"reflect"
	"strings"
)

type generic[T model.Model] struct {
	database database.Database
	model    T
}

func Generic[T model.Model](database database.Database, model T) *generic[T] {
	database.AutoMigrate(model)
	return &generic[T]{database: database, model: model}
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

func (r *generic[T]) All() []T {
	var result []T
	r.database.Find(&result)
	return result
}

func (r *generic[T]) Find(id int64) *T {
	var result []T
	r.database.Find(&result, id)
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
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

func (r *generic[T]) Raw(query string) []T {
	var result []T
	r.database.Raw("SELECT * FROM " + r.database.Model(r.model).Name() + " WHERE " + query).Scan(&result)
	return result
}

func (r *generic[T]) Count(query string, args ...string) int64 {
	var result int64
	r.database.Model(r.model).Where(query, args).Count(&result)
	return result
}

func (r *generic[T]) Insert(entity *T) {
	r.database.Create(entity)
}

func (r *generic[T]) Update(entity *T) {
	r.database.Save(entity)
}

func (r *generic[T]) Delete(entity *T) {
	r.database.Delete(entity)
}
