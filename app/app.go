package app

import (
	"flag"
	"genuine/app/controllers"
	"genuine/app/database"
	"genuine/app/decorators"
	"genuine/app/filters"
	"genuine/app/functions"
	"genuine/app/localizations"
	"genuine/app/models"
	"genuine/app/models/register"
	"genuine/app/repositories"
	"genuine/app/server"
	"genuine/core"
	controllers2 "genuine/core/controllers"
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
		database.Standard,
		decorators.Navigation,
		filters.Basic,
		functions.Form,
		functions.Paginate,
		functions.Table,
		localizations.All,
		repositories.Categories,
		repositories.Payments,
		repositories.Users,
		server.Webdav,
	)

	for _, model := range register.Models {
		provide(model)
	}

	flag.Parse()
}

func provide[T any](model T) {
	core.Provide(func(database database.Database) repositories2.Repository[T] {
		return repositories.Generic(database, model, "Id DESC")
	})
	core.Provide(func(repository repositories2.Repository[T]) controllers2.Controller {
		return controllers.Generic[T](repository, "/assets/"+models.Plural(model))
	})
}
