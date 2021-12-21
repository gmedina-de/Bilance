package repository

import (
	"Bilance/model"
	"Bilance/service"
)

type userRepository struct {
	gRepository[model.User]
}

func UserRepository(database service.Database) GRepository[model.User] {
	return &userRepository{gRepository[model.User]{database, model.User{}}}
}
