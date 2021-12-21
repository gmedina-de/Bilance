package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"net/http"
)

type projects struct {
	generic
	users repository.GRepository[model.User]
}

func Projects(repository repository.Repository, users repository.GRepository[model.User]) Controller {
	return &projects{
		generic{
			repository: repository,
			basePath:   "/admin/projects/",
			dataProvider: func(request *http.Request) interface{} {
				if request.URL.Query().Get("Id") != "" {
					idString := request.URL.Query().Get("Id")
					return users.List("WHERE Id NOT IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + idString + ")")
				} else {
					return users.List()
				}
			},
		},
		users,
	}
}
