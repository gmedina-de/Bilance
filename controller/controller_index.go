package controller

import (
	"Bilance/model"
	"Bilance/service"
	"net/http"
)

type indexController struct {
}

func IndexController() Controller {
	return &indexController{}
}

func (this *indexController) Routing(router service.Router) {
	router.Get("/", this.Index)
}

func (this *indexController) Index(writer http.ResponseWriter, request *http.Request) {
	selectedProjectId, hasParameter := request.URL.Query()[model.SelectedProjectIdCookie]
	if hasParameter {
		model.SetSelectedProjectId(writer, selectedProjectId[0])
		redirect(writer, request, "/")
	}

	render(writer, request, &Parameters{}, "dashboard", "index")
}
