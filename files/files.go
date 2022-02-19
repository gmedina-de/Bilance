package files

import (
	"genuine/core/inject"
	"genuine/core/template"
	controllers "genuine/files/controller"
)

func init() {

	template.AddNavigation("sites", "layout")
	template.AddNavigation("tasks", "check-circle")

	template.AddNavigation("files", "folder").
		WithChild("all", "folder").
		WithChild("favorites", "star").
		WithChild("last", "clock").
		WithChild("trash", "trash")

	inject.Implementations(controllers.Files)

}
