package dashboard

import (
	"homecloud/core"
	"homecloud/dashboard/controller"
)

func init() {
	core.AddConstructors(
		controller.Index,
	)
}
