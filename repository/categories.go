package repository

import (
	"Bilance/model"
	"Bilance/service"
)

type categories struct {
	generic[model.Category]
}

func Categories(database service.Database) GRepository[model.Category] {
	return &categories{generic[model.Category]{database, model.Category{}}}
}
