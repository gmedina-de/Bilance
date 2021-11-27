package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
	"strconv"
)

type paymentsController struct {
	crudController
}

func PaymentController(repository repository.Repository, categoryRepository repository.Repository, userRepository repository.Repository) Controller {
	return &paymentsController{
		crudController{
			repository: repository,
			basePath:   "/",
			dataProvider: func(request *http.Request) interface{} {
				projectId := model.GetSelectedProjectIdString(request)
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

func (c *paymentsController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"edit", c.Edit)
	router.Post(c.basePath+"edit", c.Edit)
	router.Get(c.basePath+"edit/delete", c.Delete)
	router.Get(c.basePath+"changeProject", c.ChangeProject)
}

func (c *paymentsController) ChangeProject(writer http.ResponseWriter, request *http.Request) {
	selectedProjectId, _ := request.URL.Query()[model.SelectedProjectIdCookie]
	model.SetSelectedProjectId(writer, selectedProjectId[0])
	redirect(writer, request, "/")
}

func (c *paymentsController) List(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Has("search") {
		term := request.URL.Query().Get("search")
		list := c.repository.List(
			"WHERE ProjectId = "+model.GetSelectedProjectIdString(request),
			"AND (Name LIKE '%"+term+"%'",
			"OR CategoryId IN (SELECT Id FROM Category WHERE Name LIKE '%"+term+"%')",
			"OR Date LIKE '"+term+"%')",
			"ORDER BY Date",
		).([]model.Payment)
		render(
			writer,
			request,
			&Parameters{
				Model: list,
				Toast: strconv.Itoa(len(list)) + " " + localization.Translate("records_found"),
			},
			"search_results",
			"crud_table",
			c.repository.ModelNamePlural(),
		)
	} else {
		c.crudController.List(writer, request)
	}
}
