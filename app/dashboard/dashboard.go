package dashboard

import (
	"genuine/app/dashboard/controllers"
	"genuine/core"
)

func init() {
	core.Provide(controllers.Index)
}
