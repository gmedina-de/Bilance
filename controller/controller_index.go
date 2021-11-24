package controller

import (
	"Bilance/model"
	"Bilance/service"
	"net/http"
	"time"
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

	selectedProjectId, ok := request.URL.Query()[model.SelectedProjectIdCookie]
	if ok {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		http.SetCookie(writer, &http.Cookie{
			Name:    model.SelectedProjectIdCookie,
			Value:   selectedProjectId[0],
			Path:    "/",
			Expires: expiration,
		})
		redirect(writer, request, "/")
	}

	render(writer, request, "dashboard", nil, "index")
}
