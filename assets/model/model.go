package model

import (
	"homecloud/core"
	"homecloud/core/controllers"
	"homecloud/core/database"
	"homecloud/core/repositories"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	core.Implementations(
		func(database database.Database) repositories.Repository[T] {
			return repositories.Generic(database, model)
		},
	)

	core.Implementations(
		func(repository repositories.Repository[T]) controllers.Controller {
			return controllers.Generic(repository, model)
		},
	)

}
