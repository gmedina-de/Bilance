package files

import (
	"genuine/calendar/controller"
	"genuine/calendar/server"
	"genuine/core/injector"
	"genuine/core/template"
)

func init() {

	template.AddNavigation("calendar", "users").
		WithChild("all", "folder").
		WithChild("favorites", "star").
		WithChild("last", "clock").
		WithChild("trash", "trash")

	injector.Implementations(
		controllers.Calendar,
		server.Carddav,
	)
}
