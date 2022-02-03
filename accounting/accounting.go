package accounting

import (
	"homecloud/accounting/controller"
	"homecloud/accounting/repository"
	"homecloud/core"
	"homecloud/core/template"
)

func init() {
	template.AddNavigation(
		template.MenuItem("accounting", "book", "/accounting").WithSubItems(
			template.MenuItem("payments", "layers", "/accounting/payments"),
			template.MenuItem("categories", "tag", "/accounting/categories"),
			template.MenuItem("analysis", "", ""),
			template.MenuItem("balances", "activity", "/accounting/balances"),
			template.MenuItem("expenses", "", ""),
			template.MenuItem("by_period", "bar-chart-2", "/accounting/expenses/by_period"),
			template.MenuItem("by_category", "pie-chart", "/accounting/expenses/by_category"),
		),
	)
	core.Register(
		repository.Categories,
		repository.Payments,
		controller.Payments,
		controller.Categories,
		controller.Balances,
		controller.Expenses,
	)
}
