package settings

import (
	"genuine/app/settings/controllers"
	"genuine/core"
)

func init() {
	core.Provide(controllers.Users)

}
