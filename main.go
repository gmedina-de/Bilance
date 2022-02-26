package main

import (
	_ "genuine/apps/accounting"
	_ "genuine/apps/assets"
	_ "genuine/apps/calendar"
	_ "genuine/apps/contacts"
	_ "genuine/apps/dashboard"
	_ "genuine/apps/files"
	_ "genuine/apps/users"
	"genuine/core"
	"genuine/core/server"
	"genuine/core/translator"
	"genuine/l10n"
)

func main() {

	core.Invoke(func(server server.Server, translator translator.Translator) {
		translator.Translation("de", l10n.De)
		server.Serve()
	})
}
