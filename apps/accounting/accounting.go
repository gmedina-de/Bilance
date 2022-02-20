package accounting

import (
	controllers2 "genuine/apps/accounting/controllers"
	models2 "genuine/apps/accounting/models"
	repositories2 "genuine/apps/accounting/repositories"
	"genuine/core/injector"
	"genuine/core/template"
	"github.com/beego/beego/v2/client/orm"
)

func init() {

	orm.RegisterModel(
		&models2.Category{},
		&models2.Payment{},
	)

	template.AddNavigation("accounting", "book").
		WithChild("payments", "layers").
		WithChild("categories", "tag").
		WithChild("analysis", "").
		WithChild("balances", "activity").
		WithChild("expenses", "").
		WithChild("expenses/by_period", "bar-chart-2").
		WithChild("expenses/by_category", "pie-chart").
		Path = "/accounting/payments"

	injector.Implementations(repositories2.Payments)
	injector.Implementations(repositories2.Categories)
	injector.Implementations(controllers2.Payments, controllers2.Categories, controllers2.Balances, controllers2.Expenses)
}
