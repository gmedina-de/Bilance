package files

import (
	"genuine/app/files/controller"
	"genuine/app/files/server"
	"genuine/core"
)

func init() {
	core.Provide(controllers.Files)
	core.Provide(server.Webdav)
}
