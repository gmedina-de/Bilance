package settings

import (
	"genuine/app/settings/authenticator"
	"genuine/app/settings/controllers"
	"genuine/app/settings/repositories"
	"genuine/core"
)

func init() {
	core.Provide(controllers.Users)
	core.Provide(repositories.Users)
	core.Provide(authenticator.Basic)
}
