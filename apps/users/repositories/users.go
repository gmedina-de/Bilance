package repositories

import (
	"genuine/apps/users/models"
	"genuine/core/repositories"
)

func Users() repositories.Repository[models.User] {
	return repositories.NewGeneric(models.User{})
}
