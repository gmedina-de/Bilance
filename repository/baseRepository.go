package repository

import (
	"Bilance/service/database"
)

type baseRepository struct {
	database database.Database
}

func (r *baseRepository) Insert(model interface{}) {
	r.database.Insert(model)
}

func (r *baseRepository) Update(model interface{}) {
	r.database.Update(model)
}

func (r *baseRepository) Delete(model interface{}) {
	r.database.Delete(model)
}
