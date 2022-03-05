package app

import (
	"genuine/app/controllers"
	"genuine/app/filters"
	"genuine/app/models/register"
	"genuine/app/navigation"
	"genuine/app/repositories"
	"genuine/app/server"
	_ "genuine/app/template"
	"genuine/core"
	controllers2 "genuine/core/controllers"
	"genuine/core/database"
	"genuine/core/models"
	repositories2 "genuine/core/repositories"
)

func init() {
	core.Provide(
		controllers.Balances,
		controllers.Categories,
		controllers.Expenses,
		controllers.Files,
		controllers.Index,
		controllers.Payments,
		controllers.Search,
		controllers.Users,
		filters.Basic,
		navigation.Standard,
		repositories.Categories,
		repositories.Payments,
		repositories.Users,
		server.Webdav,
	)

	for _, model := range register.Models {
		provide(model)
	}
}

func provide[T any](model T) {
	core.Provide(func(database database.Database) repositories2.Repository[T] {
		return repositories2.Generic(database, model, "Id DESC")
	})
	core.Provide(func(repository repositories2.Repository[T]) controllers2.Controller {
		return controllers.Generic[T](repository, "/assets/"+models.Plural(model))
	})
}
