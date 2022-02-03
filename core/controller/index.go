package controller

import (
	"homecloud/core/server"
	"homecloud/core/template"
	"net/http"
)

type index struct {
}

func Index() Controller {
	return &index{}
}

func (c *index) Routing(server server.Server) {
	server.Get("", c.Index)
}

func (c *index) Index(writer http.ResponseWriter, request *http.Request) {
	template.Render(writer, request, nil, "", "core/template/index.gohtml")
}

//
//func (c *index) ChangeProject(writer http.ResponseWriter, request *http.Request) {
//	selectedProjectId, _ := request.URL.Query()[model.SelectedProjectIdCookie]
//	model.SetSelectedProjectId(writer, selectedProjectId[0])
//	controller.redirect(writer, request, "/")
//}
