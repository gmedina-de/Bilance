package repositories

import (
	"genuine/core/database"
	"genuine/core/models"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
)

type Generic[T models.Model] struct {
	Database database.Database
	Model    T
	Ordering string
}

func (r Generic[T]) Init() Repository[T] {
	if r.Ordering == "" {
		r.Ordering = "Id DESC"
	}
	orm.RegisterModel(&r.Model)
	return &r
}

func (r *Generic[T]) All() []T {
	var result []T
	r.Database.Raw("SELECT * FROM " + r.modelName() + " ORDER BY " + r.Ordering).QueryRows(&result)
	return result
}

func (r *Generic[T]) Count() int64 {
	var count int64
	r.Database.Raw("SELECT COUNT(*) FROM " + r.modelName()).QueryRow(&count)
	return count
}

func (r *Generic[T]) Limit(limit int, offset int) []T {
	var result []T
	r.Database.Raw("SELECT * FROM "+r.modelName()+" ORDER BY "+r.Ordering+" LIMIT ? OFFSET ?", limit, offset).QueryRows(&result)
	return result
}

func (r *Generic[T]) Find(id int64) *T {
	var result T
	r.Database.Raw("SELECT * FROM "+r.modelName()+" WHERE Id = ?", id).QueryRow(&result)
	return &result
}

func (r *Generic[T]) List(query string, args ...any) []T {
	var result []T
	r.Database.Raw("SELECT * FROM "+r.modelName()+" "+query, args...).QueryRows(&result)
	return result
}

func (r *Generic[T]) Map(query string, args ...any) map[int64]*T {
	var result = make(map[int64]*T)
	list := r.List(query, args...)
	for _, elem := range list {
		result[reflect.ValueOf(elem).FieldByName("Id").Interface().(int64)] = &elem
	}
	return result
}

func (r *Generic[T]) Insert(entity any) {
	r.Database.Insert(entity)
}

func (r *Generic[T]) Update(entity any) {
	r.Database.Update(entity)
}

func (r *Generic[T]) Delete(entity any) {
	r.Database.Delete(entity)
}

func (r *Generic[T]) modelName() string {
	return models.Name(r.Model)
}
