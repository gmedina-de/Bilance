package files

import (
	"genuine/core/injector"
	"genuine/core/template"
	controllers "genuine/files/controller"
	"genuine/files/server"
)

func init() {

	template.AddNavigation("sites", "layout")
	template.AddNavigation("tasks", "check-circle")

	template.AddNavigation("files", "folder").
		WithChild("all", "folder").
		WithChild("favorites", "star").
		WithChild("last", "clock").
		WithChild("trash", "trash")

	injector.Implementations(
		controllers.Files,
		server.Webdav,
	)
}
