package controllers

import (
	"genuine/core/models"
)

func Users() Controller {
	return Generic(models.User{}, "/settings/users")
}
