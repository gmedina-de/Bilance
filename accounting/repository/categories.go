package repository

import (
	"homecloud/accounting/model"
	"homecloud/core/database"
	"homecloud/core/repository"
	"net/http"
)

type categories struct {
	repository.Repository[model.Category]
}

func Categories(database database.Database) repository.Repository[model.Category] {
	return &categories{repository.NewGeneric(database, model.Category{})}
}

func (c *categories) FromRequest(request *http.Request, id int64) *model.Category {
	return &model.Category{
		Id:    id,
		Name:  request.Form.Get("Name"),
		Color: request.Form.Get("Color"),
	}
}
