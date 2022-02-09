package files

import (
	"genuine/core/injector"
	"genuine/core/template"
	controllers "genuine/files/controller"
	"genuine/files/server"
)

func init() {
	template.AddNavigation("files", "file")

	injector.Implementations(
		controllers.Files,
		server.Webdav,
	)
}
