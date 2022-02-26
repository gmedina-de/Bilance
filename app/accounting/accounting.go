package accounting

import (
	controllers2 "genuine/app/accounting/controllers"
	repositories2 "genuine/app/accounting/repositories"
	"genuine/core"
)

func init() {
	core.Provide(repositories2.Payments, repositories2.Categories)
	core.Provide(controllers2.Payments, controllers2.Categories, controllers2.Balances, controllers2.Expenses)
}
