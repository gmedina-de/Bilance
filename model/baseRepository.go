package model

import (
	"Bilance/service/database"
)

type baseRepository struct {
	database database.Database
}

func (r *baseRepository) Insert(entity interface{}) {
	r.database.Insert(entity)
}

func (r *baseRepository) Update(entity interface{}) {
	r.database.Update(entity)
}

func (r *baseRepository) Delete(entity interface{}) {
	r.database.Delete(entity)
}
