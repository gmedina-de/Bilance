package users

import (
	"genuine/apps/users/authenticator"
	"genuine/apps/users/controllers"
	"genuine/apps/users/repositories"
	"genuine/core"
	"genuine/core/template"
)

func init() {
	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/settings/users"

	core.Provide(controllers.Users)
	core.Provide(repositories.Users)
	core.Provide(authenticator.Basic)
}
