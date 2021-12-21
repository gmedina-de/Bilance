package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
)

type projectsController struct {
	crudController
	userRepository repository.GRepository[model.User]
}

func ProjectsController(repository repository.Repository, userRepository repository.GRepository[model.User]) Controller {
	return &projectsController{
		crudController{
			repository: repository,
			basePath:   "/admin/projects/",
			dataProvider: func(request *http.Request) interface{} {
				if request.URL.Query().Get("Id") != "" {
					idString := request.URL.Query().Get("Id")
					return userRepository.List("WHERE Id NOT IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + idString + ")")
				} else {
					return userRepository.List()
				}
			},
		},
		userRepository,
	}
}

func (c *projectsController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"edit", c.Edit)
	router.Post(c.basePath+"edit", c.Edit)
	router.Get(c.basePath+"edit/delete", c.Delete)
}
