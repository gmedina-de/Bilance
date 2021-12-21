package controller

import (
	"Bilance/model"
	"Bilance/service"
	"net/http"
)

type index struct {
}

func Index() Controller {
	return &index{}
}

func (c *index) Routing(router service.Router) {
	router.Get("/", c.Index)
	router.Get("/changeProject", c.ChangeProject)
}

func (c *index) Index(writer http.ResponseWriter, request *http.Request) {
	render(writer, request, nil, "", "index")
}

func (c *index) ChangeProject(writer http.ResponseWriter, request *http.Request) {
	selectedProjectId, _ := request.URL.Query()[model.SelectedProjectIdCookie]
	model.SetSelectedProjectId(writer, selectedProjectId[0])
	redirect(writer, request, "/")
}
