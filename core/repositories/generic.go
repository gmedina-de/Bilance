package repositories

import (
	"genuine/core/database"
	"genuine/core/models"
)

type generic[T any] struct {
	database database.Database
	model    T
	ordering string
}

func Generic[T any](database database.Database, model T, ordering string) Repository[T] {
	database.Migrate(model)
	if ordering == "" {
		ordering = "Id DESC"
	}
	return &generic[T]{database, model, ordering}
}

func (g *generic[T]) All() []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.tableName()+" ORDER BY "+g.ordering)
	return result
}

func (g *generic[T]) Count() int64 {
	var result int64
	g.database.Select(&result, "COUNT(*) FROM "+g.tableName())
	return result
}

func (g *generic[T]) Limit(limit int64, offset int64) []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.tableName()+" ORDER BY "+g.ordering+" LIMIT ? OFFSET ?", limit, offset)
	return result
}

func (g *generic[T]) Find(id int64) *T {
	var result T
	g.database.Select(&result, "* FROM "+g.tableName()+" WHERE Id = ?", id)
	return &result
}

func (g *generic[T]) List(query string, args ...any) []T {
	var result []T
	g.database.Select(&result, "* FROM "+g.tableName()+" "+query, args...)
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

func (g *generic[T]) Model() T {
	return g.model
}

func (g *generic[T]) tableName() string {
	return models.Plural(g.model)
}
