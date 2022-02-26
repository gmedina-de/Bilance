package controllers

import (
	"genuine/app/accounting/models"
	controllers2 "genuine/app/settings/controllers"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Payments(repository repositories.Repository[models.Payment]) controllers.Controller {
	return controllers2.Generic[models.Payment](repository, "/accounting/payments")
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
