package repository

import (
	"Bilance/model"
	"Bilance/service"
)

type categoryRepository struct {
	gRepository[model.Category]
}

func CategoryRepository(database service.Database) GRepository[model.Category] {
	return &categoryRepository{gRepository[model.Category]{database, model.Category{}}}
}
