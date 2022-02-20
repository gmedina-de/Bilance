package assets

import (
	"genuine/apps/dashboard/controllers"
	"genuine/core"
	"genuine/core/template"
)

func init() {
	template.AddNavigation("home", "home").
		Path = "/"

	core.Implementations(controllers.Index)
}
