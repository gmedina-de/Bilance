package controllers

import (
	"homecloud/core/models"
	"homecloud/core/repositories"
)

func Users(repository repositories.Repository[models.User]) Controller {
	return Generic(repository, models.User{}, "/settings/users")
}
