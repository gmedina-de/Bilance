package accounting

import (
	"homecloud/accounting/controller"
	"homecloud/accounting/repository"
	"homecloud/core"
)

func init() {
	core.Register(
		repository.Categories,
		repository.Payments,
		controller.Payments,
		controller.Categories,
		controller.Balances,
		controller.Expenses,
	)
}
