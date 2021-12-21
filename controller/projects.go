package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"net/http"
)

type projects struct {
	generic[model.Project]
	users repository.Repository[model.User]
}

func Projects(repository repository.Repository[model.Project], users repository.Repository[model.User]) Controller {
	return &projects{
		generic[model.Project]{
			repository: repository,
			basePath:   "/admin/projects/",
			dataProvider: func(request *http.Request) interface{} {
				return users.List()
			},
		},
		users,
	}
}
