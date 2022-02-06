package repositories

import (
	"github.com/beego/beego/v2/client/orm"
	"homecloud/core/database"
	"homecloud/core/model"
	"reflect"
)

type generic[T model.Model] struct {
	database database.Database
	model    T
}

func Generic[T model.Model](database database.Database, model T) Repository[T] {
	orm.RegisterModel(&model)
	return &generic[T]{database: database, model: model}
}

func (r *generic[T]) All() []T {
	var result []T
	r.database.Raw("SELECT * FROM " + model.Name(r.model)).QueryRows(&result)
	return result
}
func (r *generic[T]) Count() int {
	var count int
	r.database.Raw("SELECT COUNT(*) FROM " + model.Name(r.model)).QueryRow(&count)
	return count
}
func (r *generic[T]) Limit(limit int, offset int) []T {
	var result []T
	r.database.Raw("SELECT * FROM "+model.Name(r.model)+" LIMIT ? OFFSET ?", limit, offset).QueryRows(&result)
	return result
}

func (r *generic[T]) Find(id int64) *T {
	var result T
	r.database.Raw("SELECT * FROM "+model.Name(r.model)+" WHERE Id = ?", id).QueryRow(&result)
	return &result
}

func (r *generic[T]) List(query string, args ...any) []T {
	var result []T
	r.database.Raw("SELECT * FROM "+model.Name(r.model)+" "+query, args...).QueryRows(&result)
	return result
}

func (r *generic[T]) Map(query string, args ...any) map[int64]*T {
	var result = make(map[int64]*T)
	list := r.List(query, args...)
	for _, elem := range list {
		result[reflect.ValueOf(elem).FieldByName("Id").Interface().(int64)] = &elem
	}
	return result
}

func (r *generic[T]) Insert(entity any) {
	r.database.Insert(entity)
}

func (r *generic[T]) Update(entity any) {
	r.database.Update(entity)
}

func (r *generic[T]) Delete(entity any) {
	r.database.Delete(entity)
}
