package app

import (
	"flag"
	"genuine/app/controllers"
	"genuine/app/database"
	"genuine/app/decorators"
	"genuine/app/filters"
	"genuine/app/functions"
	"genuine/app/localizations"
	"genuine/app/repositories"
	"genuine/app/server"
	"genuine/core"
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
		controllers.Sites,
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
		repositories.Sites,
		repositories.Users,
		server.Webdav,
	)

	flag.Parse()
}
