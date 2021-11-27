package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
)

type projectsController struct {
	crudController
	userRepository repository.Repository
}

func ProjectController(repository repository.Repository, userRepository repository.Repository) Controller {
	return &projectsController{
		crudController{
			repository: repository,
			basePath:   "/admin/projects/",
			dataProvider: func(request *http.Request) interface{} {
				if request.URL.Query().Has("Id") {
					idString := request.URL.Query().Get("Id")
					return userRepository.List("WHERE Id NOT IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + idString + ")").([]model.User)
				} else {
					return userRepository.List().([]model.User)
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
