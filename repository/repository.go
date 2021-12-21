package repository

import (
	"Bilance/model"
	"net/http"
)

type Repository[T model.Model] interface {
	ModelName() string
	ModelNamePlural() string
	NewEmpty() *T
	FromRequest(request *http.Request, id int64) *T
	Find(id int64) *T
	List(conditions ...string) []T
	Map(conditions ...string) map[int64]*T
	Count(conditions ...string) int64
	Insert(entity *T)
	Update(entity *T)
	Delete(entity *T)
}
