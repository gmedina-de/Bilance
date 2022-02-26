package repositories

import (
	"genuine/apps/users/models"
	"genuine/core/database"
	"genuine/core/repositories"
)

func Users(database database.Database) repositories.Repository[models.User] {
	return repositories.Generic(database, models.User{}, "Id DESC")
}
