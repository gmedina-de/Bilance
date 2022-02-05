package repository

import (
	"homecloud/core/model"
)

type Repository[T model.Model] interface {
	All() []T
	Find(id int64) *T
	List(query string, args ...string) []T
	Map(query string, args ...string) map[int64]*T

	Insert(entity any)
	Update(entity any)
	Delete(entity any)
}
