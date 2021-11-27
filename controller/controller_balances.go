package controller

import (
	"Bilance/repository"
	"Bilance/service"
	"net/http"
)

type balancesController struct {
	paymentRepository repository.Repository
}

func BalancesController(paymentRepository repository.Repository) Controller {
	return &balancesController{paymentRepository}
}

func (c *balancesController) Routing(router service.Router) {
	router.Get("/balances/", c.Balances)
}

func (c *balancesController) Balances(writer http.ResponseWriter, request *http.Request) {
	render(
		writer,
		request,
		&Parameters{},
		"balances",
		"balances",
	)
}
