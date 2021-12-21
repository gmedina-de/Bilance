package repository

import (
	"Bilance/database"
	"Bilance/model"
	"database/sql"
	"net/http"
)

type categories struct {
	generic[model.Category]
}

func Categories(database database.Database) Repository[model.Category] {
	return &categories{
		generic[model.Category]{
			database,
			model.Category{},
		},
	}
}

func (c *categories) FromQuery(row *sql.Rows) *model.Category {
	category := model.Category{}
	model.ScanAndPanic(row, &category.Id, &category.Name, &category.Color, &category.ProjectId)
	return &category
}

func (c *categories) FromRequest(request *http.Request, id int64) *model.Category {
	return &model.Category{
		Id:        id,
		Name:      request.Form.Get("Name"),
		Color:     request.Form.Get("Color"),
		ProjectId: model.GetSelectedProjectId(request),
	}
}
