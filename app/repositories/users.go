package repositories

import (
	"genuine/app/database"
	"genuine/app/models"
	"genuine/core/repositories"
)

func Users(database database.Database) repositories.Repository[models.User] {
	r := Generic(database, models.User{}, "Id DESC")
	if r.Count("is_admin = ?", true) == 0 {
		r.Insert(&models.User{
			Name:     "admin",
			Password: "admin",
			IsAdmin:  true,
		})
	}
	return r
}
