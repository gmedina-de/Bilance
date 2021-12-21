package repository

import (
	"Bilance/model"
	"Bilance/service"
)

type users struct {
	generic[model.User]
}

func Users(database service.Database) GRepository[model.User] {
	return &users{generic[model.User]{database, model.User{}}}
}
