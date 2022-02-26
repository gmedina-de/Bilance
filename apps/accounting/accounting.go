package accounting

import (
	controllers2 "genuine/apps/accounting/controllers"
	repositories2 "genuine/apps/accounting/repositories"
	"genuine/core"
	"genuine/core/template"
)

func init() {

	template.AddNavigation("accounting", "book").
		WithChild("payments", "layers").
		WithChild("categories", "tag").
		WithChild("analysis", "").
		WithChild("balances", "activity").
		WithChild("expenses", "").
		WithChild("expenses/by_period", "bar-chart-2").
		WithChild("expenses/by_category", "pie-chart").
		Path = "/accounting/payments"

	core.Provide(repositories2.Payments, repositories2.Categories)
	core.Provide(controllers2.Payments, controllers2.Categories, controllers2.Balances, controllers2.Expenses)
}
