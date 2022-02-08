package controllers

import (
	models3 "genuine/accounting/models"
	"genuine/core/controllers"
	"genuine/core/models"
	"genuine/core/repositories"
)

func Payments(
	repository repositories.Repository[models3.Payment],
	categories repositories.Repository[models3.Category],
	users repositories.Repository[models.User],
) controllers.Controller {
	return controllers.Generic(repository, models3.Payment{}, "/accounting/payments")
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
