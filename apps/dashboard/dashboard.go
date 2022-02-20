package assets

import (
	"genuine/apps/dashboard/controllers"
	"genuine/core/injector"
	"genuine/core/template"
)

func init() {
	template.AddNavigation("home", "home").
		Path = "/"

	injector.Implementations(controllers.Index)
}
