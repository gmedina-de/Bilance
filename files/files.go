package files

import (
	"genuine/core/injector"
	"genuine/core/template"
	controllers "genuine/files/controller"
	"genuine/files/server"
)

func init() {

	template.AddNavigation("e-mail", "mail")
	template.AddNavigation("contacts", "users")
	template.AddNavigation("calendar", "calendar")
	template.AddNavigation("passwords", "key")
	template.AddNavigation("tasks", "check-circle")

	template.AddNavigation("files", "file").
		WithChild("all", "folder").
		WithChild("favorites", "star").
		WithChild("last", "clock").
		WithChild("trash", "trash")

	injector.Implementations(
		controllers.Files,
		server.Webdav,
	)
}
