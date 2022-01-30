package repository

import (
	"Bilance/database"
	"Bilance/model"
	"net/http"
)

type projects struct {
	*generic[model.Project]
}

func Projects(database database.Database) Repository[model.Project] {
	return &projects{Generic(database, model.Project{})}

}

func (pr projects) FromRequest(request *http.Request, id int64) *model.Project {
	return &model.Project{
		Id:          id,
		Name:        request.Form.Get("Name"),
		Description: request.Form.Get("Description"),
		UserIds:     request.Form.Get("UserIds"),
	}
}
