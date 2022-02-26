package main

import (
	_ "genuine/app/accounting"
	_ "genuine/app/assets"
	"genuine/app/assets/models"
	_ "genuine/app/calendar"
	_ "genuine/app/contacts"
	_ "genuine/app/dashboard"
	_ "genuine/app/files"
	_ "genuine/app/settings"
	"genuine/core"
	"genuine/core/localization"
	model2 "genuine/core/models"
	"genuine/core/navigation"
	"genuine/core/server"
	"genuine/l10n"
)

func main() {

	core.Invoke(func(server server.Server, nav navigation.Navigation, translator localization.Localization) {
		fill(nav)
		translator.Add("de", l10n.De)
		server.Serve()
	})
}

func fill(nav navigation.Navigation) {
	nav.Add("home", "home").Path = "/"
	nav.Add("accounting", "book").
		Sub("payments", "layers").
		Sub("categories", "tag").
		Sub("analysis", "").
		Sub("balances", "activity").
		Sub("expenses", "").
		Sub("expenses/by_period", "bar-chart-2").
		Sub("expenses/by_category", "pie-chart").
		Path = "/accounting/payments"

	menuItem := nav.Add("assets", "box")
	Models := models.Models
	for i, m := range Models {
		menuItem = menuItem.Sub(model2.Plural(m), models.Icons[i])
	}
	menuItem.Path = "/assets/" + model2.Plural(Models[0])

	nav.Add("files", "folder").
		Sub("all", "folder").
		Sub("favorites", "star").
		Sub("last", "clock").
		Sub("trash", "trash")

	nav.Add("sites", "layout")
	nav.Add("tasks", "check-circle")
	nav.Add("settings", "settings").
		Sub("users", "users").
		Path = "/settings/users"
}
