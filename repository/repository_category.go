package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
)

type categoryRepository struct {
	genericRepository[model.Category]
}

func CategoryRepository(database service.Database) GenericRepository[model.Category] {
	return &categoryRepository{
		genericRepository[model.Category]{
			database,
			model.Category{},
			func(row *sql.Rows) *model.Category {
				var id int64
				var name string
				var color string
				var projectId int64
				scanAndPanic(row, &id, &name, &color, &projectId)
				return &model.Category{id, name, color, projectId}
			},
			func(request *http.Request, id int64) *model.Category {
				projectId := model.GetSelectedProjectId(request)
				return &model.Category{
					id,
					request.Form.Get("Name"),
					request.Form.Get("Color"),
					projectId,
				}
			},
		},
	}
}
