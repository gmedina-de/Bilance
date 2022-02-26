package files

import (
	"genuine/app/files/controller"
	"genuine/core"
)

func init() {
	core.Provide(controllers.Files)
}
