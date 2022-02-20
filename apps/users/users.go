package users

import (
	"genuine/apps/users/authenticator"
	"genuine/apps/users/controllers"
	"genuine/core/injector"
	"genuine/core/repositories"
)

func init() {
	injector.Implementations(controllers.Users)
	injector.Implementations(repositories.Users)
	injector.Implementations(authenticator.Basic)
}
