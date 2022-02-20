package repositories

import (
	"genuine/core/database"
	"genuine/core/models"
)

type Generic[T models.Model] struct {
	database database.Database
	model    T
	ordering string
}

func NewGeneric[T models.Model](model T) Repository[T] {
	return &Generic[T]{model: model, ordering: "Id DESC"}
}

func (g *Generic[T]) All() []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.modelName()+" ORDER BY "+g.ordering)
	return result
}

func (g *Generic[T]) Count() int64 {
	var result int64
	g.database.Select(&result, "COUNT(*) FROM "+g.modelName())
	return result
}

func (g *Generic[T]) Limit(limit int, offset int) []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.modelName()+" ORDER BY "+g.ordering+" LIMIT ? OFFSET ?", limit, offset)
	return result
}

func (g *Generic[T]) Find(id int64) *T {
	var result T
	g.database.Select(&result, "* FROM "+g.modelName()+" WHERE Id = ?", id)
	return &result
}

func (g *Generic[T]) List(query string, args ...any) []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.modelName()+" "+query, args...)
	return result
}

func (g *Generic[T]) Insert(entity any) {
	g.database.Insert(entity)
}

func (g *Generic[T]) Update(entity any) {
	g.database.Update(entity)
}

func (g *Generic[T]) Delete(entity any) {
	g.database.Delete(entity)
}

func (g *Generic[T]) T() T {
	return g.model
}

func (g *Generic[T]) modelName() string {
	return models.Name(g.model)
}
