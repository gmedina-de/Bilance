package main

import (
	_ "genuine/accounting"
	_ "genuine/assets"
	"genuine/core"
	_ "genuine/files"
)

func main() {
	core.Init()
}
