package repositories

import (
	"genuine/app/common/models"
	"genuine/core/database"
	"genuine/core/repositories"
)

func Users(database database.Database) repositories.Repository[models.User] {
	return repositories.Generic(database, models.User{}, "Id DESC")
}
