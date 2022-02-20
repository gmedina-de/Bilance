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
	"genuine/core/template"
	"genuine/core/translator"
	"genuine/l10n"
)

func main() {
	core.Init(App).Run()
}

type app struct {
	server     server.Server
	Translator translator.Translator
	Template   template.Template
}

func App() *app {
	return &app{}
}

func (a *app) Run() {
	a.Template.Parse("views")
	a.Translator.Translation("de", l10n.De)
	a.server.Start()
}
