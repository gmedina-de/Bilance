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
)

func main() {
	core.Init()
}
