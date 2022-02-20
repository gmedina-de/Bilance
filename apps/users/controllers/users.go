package controllers

import (
	"genuine/apps/users/models"
	"genuine/core/controllers"
)

func Users() controllers.Controller {
	return controllers.Generic[models.User]("/settings/users")
}
