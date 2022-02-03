package controller

import (
	"homecloud/core/controller"
	"net/http"
)

type index struct {
}

func Index() controller.Controller {
	return &index{}
}

func (c *index) Index(writer http.ResponseWriter, request *http.Request) {
	controller.Render(writer, request, nil, "", "index")
}

//
//func (c *index) ChangeProject(writer http.ResponseWriter, request *http.Request) {
//	selectedProjectId, _ := request.URL.Query()[model.SelectedProjectIdCookie]
//	model.SetSelectedProjectId(writer, selectedProjectId[0])
//	controller.redirect(writer, request, "/")
//}
