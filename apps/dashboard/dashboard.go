package dashboard

import (
	"genuine/apps/dashboard/controllers"
	"genuine/core"
	"genuine/core/template"
)

func init() {
	template.AddNavigation("home", "home").
		Path = "/"

	core.Provide(controllers.Index)
}
