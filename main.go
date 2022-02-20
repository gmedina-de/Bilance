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
	"genuine/core/injector"
	"genuine/core/server"
	"genuine/core/translator"
	"genuine/l10n"
)

func main() {
	core.Init(App)
}

type app struct {
	Server     server.Server
	Translator translator.Translator
}

func App() injector.Initiable {
	return &app{}
}

func (a *app) Init() {
	a.Translator.Translation("de", l10n.De)
	a.Server.Start()
}
