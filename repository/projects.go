package repository

import (
	"Bilance/database"
	"Bilance/model"
	"database/sql"
	"net/http"
)

type projects struct {
	generic[model.Project]
}

func Projects(database database.Database) Repository[model.Project] {
	return &projects{
		generic[model.Project]{
			database: database,
			model:    model.Project{},
		},
	}
}

func (pr projects) FromQuery(row *sql.Rows) *model.Project {
	project := model.Project{}
	model.ScanAndPanic(row, &project.Id, &project.Name, &project.Description, &project.UserIds)
	return &project
}

func (pr projects) FromRequest(request *http.Request, id int64) *model.Project {
	return &model.Project{
		Id:          id,
		Name:        request.Form.Get("Name"),
		Description: request.Form.Get("Description"),
		UserIds:     request.Form.Get("UserIds"),
	}
}
