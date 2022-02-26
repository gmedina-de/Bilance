package repositories

import (
	"genuine/core/database"
	"genuine/core/models"
)

type generic[T models.Model] struct {
	database database.Database
	model    T
	ordering string
}

func Generic[T models.Model](database database.Database, model T, ordering string) Repository[T] {
	database.Migrate(model)
	if ordering == "" {
		ordering = "Id DESC"
	}
	return &generic[T]{database, model, ordering}
}

func (g *generic[T]) All() []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.modelName()+" ORDER BY "+g.ordering)
	return result
}

func (g *generic[T]) Count() int64 {
	var result int64
	g.database.Select(&result, "COUNT(*) FROM "+g.modelName())
	return result
}

func (g *generic[T]) Limit(limit int, offset int) []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.modelName()+" ORDER BY "+g.ordering+" LIMIT ? OFFSET ?", limit, offset)
	return result
}

func (g *generic[T]) Find(id int64) *T {
	var result T
	g.database.Select(&result, "* FROM "+g.modelName()+" WHERE Id = ?", id)
	return &result
}

func (g *generic[T]) List(query string, args ...any) []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.modelName()+" "+query, args...)
	return result
}

func (g *generic[T]) Insert(entity any) {
	g.database.Insert(entity)
}

func (g *generic[T]) Update(entity any) {
	g.database.Update(entity)
}

func (g *generic[T]) Delete(entity any) {
	g.database.Delete(entity)
}

func (g *generic[T]) T() T {
	return g.model
}

func (g *generic[T]) modelName() string {
	return models.Plural(g.model)
}
