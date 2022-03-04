package repositories

import (
	"genuine/core/database"

	"gorm.io/gorm/clause"
)

type generic[T any] struct {
	database database.Database
	model    T
	ordering string
}

func Generic[T any](database database.Database, model T, ordering string) Repository[T] {
	if ordering == "" {
		ordering = "Id DESC"
	}

	err := database.AutoMigrate(&model)
	if err != nil {
		database.Logger.Error(nil, err.Error())
	}

	return &generic[T]{database, model, ordering}
}

func (g *generic[T]) All() []T {
	var result []T
	g.database.Preload(clause.Associations).Find(&result)
	return result
}

func (g *generic[T]) Count() int64 {
	var result int64
	g.database.Model(&g.model).Count(&result)
	return result
}

func (g *generic[T]) Limit(limit int, offset int) []T {
	var result []T
	g.database.Limit(limit).Offset(offset).Preload(clause.Associations).Find(&result)
	return result
}

func (g *generic[T]) Find(id uint) *T {
	var result T
	g.database.Preload(clause.Associations).First(&result, id)
	return &result
}

func (g *generic[T]) List(where string, args ...any) []T {
	var result []T
	g.database.Where(where, args...).Find(&result)
	return result
}

func (g *generic[T]) Insert(entity *T) {
	g.database.Create(entity)
}

func (g *generic[T]) Update(entity *T) {
	g.database.Save(entity)
}

func (g *generic[T]) Delete(entity *T) {
	g.database.Delete(entity)
}

func (g *generic[T]) Model() T {
	return g.model
}
