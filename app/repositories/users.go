package repositories

import (
	"genuine/app/database"
	"genuine/app/models"
	"genuine/core/repositories"
)

func Users(database database.Database) repositories.Repository[models.User] {
	return Generic(database, models.User{}, "Id DESC")
}
