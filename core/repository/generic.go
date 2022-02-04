package repository

import (
	"homecloud/core/database"
	"homecloud/core/model"
	"net/http"
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

func (r *Generic[T]) NewEmpty() *T {
	return &r.model
}

func (r *Generic[T]) FromRequest(request *http.Request, id int64) *T {
	//TODO implement me
	panic("implement me")
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
		return r.NewEmpty()
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

func (r *Generic[T]) Raw(query string) []T {
	var result []T
	r.database.Raw("SELECT * FROM " + r.database.Model(r.model).Name() + " WHERE " + query).Scan(&result)
	return result
}

func (r *Generic[T]) Count(query string, args ...string) int64 {
	var result int64
	r.database.Model(r.model).Where(query, args).Count(&result)
	return result
}

func (r *Generic[T]) Insert(entity *T) {
	r.database.Create(entity)
}

func (r *Generic[T]) Update(entity *T) {
	r.database.Save(entity)
}

func (r *Generic[T]) Delete(entity *T) {
	r.database.Delete(entity)
}
