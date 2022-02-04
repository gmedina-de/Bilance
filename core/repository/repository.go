package repository

import (
	"homecloud/core/model"
	"net/http"
)

type Repository[T model.Model] interface {
	NewEmpty() *T
	FromRequest(request *http.Request, id int64) *T

	All() []T
	Find(id int64) *T
	List(query string, args ...string) []T
	Map(query string, args ...string) map[int64]*T
	Raw(sql string) []T

	Count(query string, args ...string) int64

	Insert(entity *T)
	Update(entity *T)
	Delete(entity *T)
}
