package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
	"strconv"
)

type payments struct {
	generic2[model.Payment]
}

func Payments(
	repository repository.GRepository[model.Payment],
	categories repository.GRepository[model.Category],
	users repository.GRepository[model.User],
) Controller {
	return &payments{
		generic2[model.Payment]{
			repository: repository,
			basePath:   "/payments/",
			dataProvider: func(request *http.Request) interface{} {
				projectId := model.GetSelectedProjectIdString(request)
				return struct {
					Categories []model.Category
					Users      []model.User
				}{
					categories.List("WHERE ProjectId = " + projectId),
					users.List("WHERE Id IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + projectId + ")"),
				}
			},
		},
	}
}

func (c *payments) Routing(router service.Router) {
	c.generic2.Routing(router)
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
}

func (c *payments) List(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("search") != "" {
		term := request.URL.Query().Get("search")
		list := c.repository.List(
			"WHERE ProjectId = "+model.GetSelectedProjectIdString(request),
			"AND (Name LIKE '%"+term+"%'",
			"OR CategoryId IN (SELECT Id FROM Category WHERE Name LIKE '%"+term+"%')",
			"OR Date LIKE '"+term+"%')",
			"ORDER BY Date",
		)
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
		c.generic2.List(writer, request)
	}
}
