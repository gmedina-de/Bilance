package models

import (
	"genuine/framework"
	"genuine/framework/controllers"
	"genuine/framework/database"
	model2 "genuine/framework/models"
	"genuine/framework/repositories"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	framework.Implementations(
		func(database database.Database) repositories.Repository[T] {
			return repositories.Generic(database, model)
		},
	)

	framework.Implementations(
		func(repository repositories.Repository[T]) controllers.Controller {
			return controllers.Generic(repository, model, "/assets/"+model2.Plural(model))
		},
	)

}
