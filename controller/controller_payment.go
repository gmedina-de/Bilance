package controller

import (
	"Bilance/repository"
	"Bilance/service"
)

type paymentController struct {
	baseController
}

func PaymentController(repository repository.Repository) Controller {
	return &paymentController{
		baseController{
			repository: repository,
			basePath:   "/payments",
		},
	}
}

func (c *paymentController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"/edit", c.Edit)
	router.Post(c.basePath+"/edit", c.Edit)
	router.Get(c.basePath+"/edit/delete", c.Delete)
}
