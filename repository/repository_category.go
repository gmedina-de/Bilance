package repository

import (
	"Bilance/model"
	"Bilance/service"
)

type categoryRepository struct {
	GenericRepository[model.Category]
}

func CategoryRepository(database service.Database) GRepository[model.Category] {
	return &categoryRepository{
		GenericRepository[model.Category]{
			database,
			model.Category{},
		},
	}
}
