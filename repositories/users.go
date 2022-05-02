package repositories

import (
	"genuine/database"
	"genuine/models"
)

func Users(database database.Database) Repository[models.User] {
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
