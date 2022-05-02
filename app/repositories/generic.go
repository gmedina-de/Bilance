package repositories

import (
	"genuine/app/database"
	"genuine/core/repositories"
	"gorm.io/gorm/clause"
)

type generic[T any] struct {
	database database.Database
	model    T
	ordering string
}

func Generic[T any](database database.Database, model T, ordering string) repositories.Repository[T] {
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
	g.database.Preload(clause.Associations).Order(g.ordering).Find(&result)
	return result
}

func (g *generic[T]) Count(where string, args ...any) int64 {
	var result int64
	g.database.Model(&g.model).Where(where, args...).Count(&result)
	return result
}

func (g *generic[T]) Limit(limit int, offset int, where string, args ...any) []T {
	var result []T
	g.database.Where(where, args...).Order(g.ordering).Limit(limit).Offset(offset).Preload(clause.Associations).Find(&result)
	return result
}

func (g *generic[T]) Find(id uint) *T {
	var result T
	g.database.Preload(clause.Associations).First(&result, id)
	return &result
}

func (g *generic[T]) List(where string, args ...any) []T {
	var result []T
	g.database.Order(g.ordering).Where(where, args...).Preload(clause.Associations).Find(&result)
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
