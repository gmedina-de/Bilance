package main

import (
	_ "genuine/accounting"
	_ "genuine/assets"
	_ "genuine/contacts"
	"genuine/core"
	_ "genuine/files"
)

func main() {
	core.Init()
}
