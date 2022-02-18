package accounting

import (
	"genuine/accounting/controllers"
	"genuine/accounting/models"
	"genuine/accounting/repositories"
	"genuine/core/injector"
	"genuine/core/template"
	"github.com/beego/beego/v2/client/orm"
)

func init() {

	orm.RegisterModel(
		&models.Category{},
		&models.Payment{},
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

	injector.Implementations(repositories.Payments)
	injector.Implementations(repositories.Categories)
	injector.Implementations(controllers.Payments, controllers.Categories, controllers.Balances, controllers.Expenses)
}
