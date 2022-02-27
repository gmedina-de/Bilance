package main

import (
	_ "genuine/app/accounting"
	_ "genuine/app/assets"
	_ "genuine/app/calendar"
	_ "genuine/app/common"
	_ "genuine/app/contacts"
	_ "genuine/app/dashboard"
	_ "genuine/app/files"
	_ "genuine/app/settings"
	"genuine/core"
	"genuine/core/server"
	"genuine/core/template"
	"genuine/core/translator"
	"genuine/l10n"
)

func main() {
	core.Invoke(func(server server.Server, template template.Template, translator translator.Translator) {
		translator.Add("de", l10n.De)
		server.Serve()
	})
}
