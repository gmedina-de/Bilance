package repositories

import (
	"homecloud/core/models"
)

type Repository[T models.Model] interface {
	All() []T
	Count() int
	Find(id int64) *T
	Limit(limit int, offset int) []T
	List(query string, args ...any) []T
	Map(query string, args ...any) map[int64]*T

	Insert(entity any)
	Update(entity any)
	Delete(entity any)
}
