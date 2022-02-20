package files

import (
	"genuine/apps/files/controller"
	"genuine/core"
	"genuine/core/template"
)

func init() {

	template.AddNavigation("sites", "layout")
	template.AddNavigation("tasks", "check-circle")

	template.AddNavigation("files", "folder").
		WithChild("all", "folder").
		WithChild("favorites", "star").
		WithChild("last", "clock").
		WithChild("trash", "trash")

	core.Implementations(controllers.Files)

}
