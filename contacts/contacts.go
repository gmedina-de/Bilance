package files

import (
	"genuine/contacts/controller"
	"genuine/contacts/server"
	"genuine/core/injector"
	"genuine/core/template"
)

func init() {

	template.AddNavigation("contacts", "users").
		WithChild("all", "folder").
		WithChild("favorites", "star").
		WithChild("last", "clock").
		WithChild("trash", "trash")

	injector.Implementations(
		controllers.Contacts,
		server.Carddav,
	)
}
