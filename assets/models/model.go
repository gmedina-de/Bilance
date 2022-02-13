package models

import (
	"genuine/core/controllers"
	"genuine/core/database"
	"genuine/core/injector"
	model2 "genuine/core/models"
	"genuine/core/repositories"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	injector.Implementations(
		func(database database.Database) repositories.Repository[T] {
			return &repositories.Generic[T]{
				Database: database,
				Model:    model,
			}
		},
	)

	injector.Implementations(
		func(repository repositories.Repository[T]) controllers.Controller {
			return controllers.Generic(repository, model, "/assets/"+model2.Plural(model))
		},
	)

}
