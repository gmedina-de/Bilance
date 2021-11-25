package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
	"strconv"
)

type projectController struct {
	baseController
	userRepository repository.Repository
}

func ProjectController(repository repository.Repository, userRepository repository.Repository) Controller {
	return &projectController{
		baseController{
			repository: repository,
			basePath:   "/admin/projects",
		},
		userRepository,
	}
}

func (c *projectController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"/edit", c.Edit)
	router.Post(c.basePath+"/edit", c.Edit)
	router.Get(c.basePath+"/edit/delete", c.Delete)
}

func (c *projectController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		if request.URL.Query().Has("Id") {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			item := c.repository.Find(id)
			notUsers := c.userRepository.List("WHERE Id NOT IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + idString + ")").([]model.User)
			render(writer, request, &Parameters{Model: item, Data: notUsers}, "edit", "crud_form", c.modelName())
		} else {
			render(writer, request, &Parameters{Model: c.repository.NewEmpty(), Data: nil}, "new", "crud_form", c.modelName())
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Has("Id") {
			id, _ := strconv.ParseInt(request.Form.Get("Id"), 10, 64)
			c.repository.Update(c.repository.NewFromRequest(request, id))
		} else {
			c.repository.Insert(c.repository.NewFromRequest(request, 0))
		}
		redirect(writer, request, c.basePath)
	}
}
