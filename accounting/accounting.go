package accounting

import (
	"homecloud/accounting/controllers"
	"homecloud/accounting/repositories"
	"homecloud/core"
	"homecloud/core/template"
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

	core.Implementations(
		repositories.Categories,
		repositories.Payments,
		controllers.Payments,
		controllers.Categories,
		controllers.Balances,
		controllers.Expenses,
	)
}
