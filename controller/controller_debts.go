package controller

import (
	"Bilance/repository"
	"Bilance/service"
	"net/http"
)

type debtsController struct {
	paymentRepository repository.Repository
}

func DebtsController(paymentRepository repository.Repository) Controller {
	return &debtsController{paymentRepository}
}

func (c *debtsController) Routing(router service.Router) {
	router.Get("/debts/", c.Debts)
}

func (c *debtsController) Debts(writer http.ResponseWriter, request *http.Request) {
	render(
		writer,
		request,
		&Parameters{},
		"debts",
		"debts",
	)
}
