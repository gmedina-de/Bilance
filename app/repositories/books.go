package repositories

import (
	"genuine/app/database"
	"genuine/app/models"
	"genuine/core/repositories"
)

func Books(database database.Database) repositories.Repository[models.Book] {
	return Generic(database, models.Book{}, "Name ASC")
}
