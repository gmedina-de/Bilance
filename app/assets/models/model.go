package models

import (
	"genuine/core"
	"genuine/core/controllers"
	"genuine/core/database"
	"genuine/core/models"
	"genuine/core/repositories"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	core.Provide(func(database database.Database) repositories.Repository[T] {
		return repositories.Generic(database, model, "Id DESC")
	})

	core.Provide(func(repository repositories.Repository[T]) controllers.Controller {
		return controllers.Generic[T](repository, "/assets/"+models.Plural(model))
	})
}
