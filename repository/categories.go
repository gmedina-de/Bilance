package repository

import (
	"Bilance/database"
	"Bilance/model"
	"net/http"
)

type categories struct {
	*generic[model.Category]
}

func Categories(database database.Database) Repository[model.Category] {
	return &categories{Generic(database, model.Category{})}
}

func (c *categories) FromRequest(request *http.Request, id int64) *model.Category {
	return &model.Category{
		Id:        id,
		Name:      request.Form.Get("Name"),
		Color:     request.Form.Get("Color"),
		ProjectId: model.GetSelectedProjectId(request),
	}
}
