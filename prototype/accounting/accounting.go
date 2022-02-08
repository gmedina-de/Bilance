package accounting

import (
	"genuine/framework"
	"genuine/framework/template"
	controllers2 "genuine/prototype/accounting/controllers"
	repositories2 "genuine/prototype/accounting/repositories"
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

	framework.Implementations(
		repositories2.Categories,
		repositories2.Payments,
		controllers2.Payments,
		controllers2.Categories,
		controllers2.Balances,
		controllers2.Expenses,
	)
}
