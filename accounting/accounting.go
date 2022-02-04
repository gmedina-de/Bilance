package accounting

import (
	"homecloud/accounting/controller"
	"homecloud/accounting/repository"
	"homecloud/core"
	"homecloud/core/template"
)

func init() {
	template.AddNavigation("accounting", "book", "/accounting/payments").
		WithChild("payments", "layers", "/accounting/payments").
		WithChild("categories", "tag", "/accounting/categories").
		WithChild("analysis", "", "").
		WithChild("balances", "activity", "/accounting/balances").
		WithChild("expenses", "", "").
		WithChild("by_period", "bar-chart-2", "/accounting/expenses/by_period").
		WithChild("by_category", "pie-chart", "/accounting/expenses/by_category")

	core.Implementations(
		repository.Categories,
		repository.Payments,
		controller.Payments,
		controller.Categories,
		controller.Balances,
		controller.Expenses,
	)
}
