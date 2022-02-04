package assets

import (
	"homecloud/assets/model"
	"homecloud/core"
	"homecloud/core/controller"
	"homecloud/core/database"
	model2 "homecloud/core/model"
	"homecloud/core/repository"
	"homecloud/core/template"
)

func init() {
	models := model.Models
	for _, m := range models {
		core.Implementations(controllerProvider(m))
		core.Implementations(repositoryProvider(m))
	}

	menuItem := template.AddNavigation("assets", "box")
	for i, m := range models {
		menuItem.WithChild(model2.Plural(m), model.Icons[i])
	}
	menuItem.Path = "/assets/" + model2.Plural(models[0])
}

func controllerProvider[T any](model T) func(repository repository.Repository[T]) controller.Controller {
	return func(repository repository.Repository[T]) controller.Controller {
		return &controller.Generic[T]{
			Repository:   repository,
			BaseTemplate: "assets/template/model.gohtml",
			BasePath:     "/assets/" + model2.Plural(model),
		}
	}
}

func repositoryProvider[T any](model T) func(database database.Database) repository.Repository[T] {
	return func(database database.Database) repository.Repository[T] {
		return repository.NewGeneric(database, model)
	}
}
