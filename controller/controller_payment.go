package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
)

type paymentController struct {
	baseController
}

func PaymentController(repository repository.Repository, categoryRepository repository.Repository, userRepository repository.Repository) Controller {
	return &paymentController{
		baseController{
			repository: repository,
			basePath:   "/payments",
			dataProvider: func(request *http.Request) interface{} {
				cookie, _ := request.Cookie(model.SelectedProjectIdCookie)
				projectId := cookie.Value
				return struct {
					Categories []model.Category
					Users      []model.User
				}{
					categoryRepository.List("WHERE ProjectId = " + projectId).([]model.Category),
					userRepository.List("WHERE Id IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + projectId + ")").([]model.User),
				}
			},
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
