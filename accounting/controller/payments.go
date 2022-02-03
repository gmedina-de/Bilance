package controller

import (
	model2 "homecloud/accounting/model"
	"homecloud/core/controller"
	"homecloud/core/localization"
	"homecloud/core/model"
	"homecloud/core/repository"
	"net/http"
	"strconv"
)

type payments struct {
	controller.Generic[model2.Payment]
}

func Payments(
	repository repository.Repository[model2.Payment],
	categories repository.Repository[model2.Category],
	users repository.Repository[model.User],
) controller.Controller {
	return &payments{
		controller.Generic[model2.Payment]{
			Repository: repository,
			DataProvider: func(request *http.Request) interface{} {
				return struct {
					Categories map[int64]*model2.Category
					Users      map[int64]*model.User
				}{
					categories.Map(""),
					users.Map(""),
				}
			},
		},
	}
}

func (c *payments) List(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("search") != "" {
		term := request.URL.Query().Get("search")
		list := c.Repository.List(
			"Name LIKE '%"+term+"%'",
			"OR CategoryId IN (SELECT Id FROM Category WHERE Name LIKE '%"+term+"%')",
			"OR Date LIKE '"+term+"%'",
			"ORDER BY Date",
		)
		controller.Render(
			writer,
			request,
			&controller.Parameters{
				Model: list,
				Toast: strconv.Itoa(len(list)) + " " + localization.Translate("records_found"),
			},
			"search_results",
			"crud_table",
			c.Repository.ModelNamePlural(),
		)
	} else {
		c.Generic.Index(writer, request)
	}
}
