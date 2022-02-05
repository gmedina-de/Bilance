package accounting

import (
	"homecloud/accounting/controller"
	"homecloud/accounting/repository"
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
		repository.Categories,
		repository.Payments,
		controller.Payments,
		controller.Categories,
		controller.Balances,
		controller.Expenses,
	)
}
