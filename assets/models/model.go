package models

import (
	"homecloud/core/controllers"
	"homecloud/core/database"
	"homecloud/core/injector"
	model2 "homecloud/core/models"
	"homecloud/core/repositories"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	injector.Implementations(
		func(database database.Database) repositories.Repository[T] {
			return repositories.Generic(database, model)
		},
	)

	injector.Implementations(
		func(repository repositories.Repository[T]) controllers.Controller {
			return controllers.Generics(repository, model, "/assets/"+model2.Plural(model))
		},
	)

}
