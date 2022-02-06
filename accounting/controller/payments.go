package controller

import (
	model2 "homecloud/accounting/model"
	"homecloud/core/controllers"
	"homecloud/core/model"
	"homecloud/core/repositories"
)

func Payments(
	repository repositories.Repository[model2.Payment],
	categories repositories.Repository[model2.Category],
	users repositories.Repository[model.User],
) controllers.Controller {
	return controllers.Generic(repository, model2.Payment{})
}

//
//func (c *payments) List(writer http.ResponseWriter, request *http.Request) {
//	if request.URL.Query().Get("search") != "" {
//		term := request.URL.Query().Get("search")
//		list := c.repository.List(
//			"Name LIKE '%"+term+"%'",
//			"OR CategoryId IN (SELECT Id FROM Category WHERE Name LIKE '%"+term+"%')",
//			"OR Date LIKE '"+term+"%'",
//			"ORDER BY Date",
//		)
//		template.Render(writer, request, "search_results", &template.Parameters{
//			model: list,
//			Toast: strconv.Itoa(len(list)) + " " + localization.Translate("records_found"),
//		}, "crud_table", "payments")
//	} else {
//		c.GenericOld.Get(writer, request)
//	}
//}
