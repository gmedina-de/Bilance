package controllers

import (
	"genuine/framework/controllers"
	"genuine/framework/models"
	"genuine/framework/repositories"
	models2 "genuine/prototype/accounting/models"
)

func Payments(
	repository repositories.Repository[models2.Payment],
	categories repositories.Repository[models2.Category],
	users repositories.Repository[models.User],
) controllers.Controller {
	return controllers.Generic(repository, models2.Payment{}, "/accounting/payments")
}

//
//func (c *payments) List(writer http.ResponseWriter, request *http.Request) {
//	if request.URL.Query().Get("search") != "" {
//		term := request.URL.Query().Get("search")
//		list := c.repositories.List(
//			"Name LIKE '%"+term+"%'",
//			"OR CategoryId IN (SELECT Id FROM Category WHERE Name LIKE '%"+term+"%')",
//			"OR Date LIKE '"+term+"%'",
//			"ORDER BY Date",
//		)
//		template.Render(writer, request, "search_results", &template.Parameters{
//			models: list,
//			Toast: strconv.Itoa(len(list)) + " " + localization.Translate("records_found"),
//		}, "crud_table", "payments")
//	} else {
//		c.GenericOld.Get(writer, request)
//	}
//}
