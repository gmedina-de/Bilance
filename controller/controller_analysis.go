package controller

import (
	"Bilance/service"
	"net/http"
)

type analysisController struct {
}

func AnalysisController() Controller {
	return &analysisController{}
}

func (this *analysisController) Routing(router service.Router) {
	router.Get("/analysis", this.Analysis)
}

func (this *analysisController) Analysis(writer http.ResponseWriter, request *http.Request) {
	render(writer, request, "analysis", nil, "analysis")
}
