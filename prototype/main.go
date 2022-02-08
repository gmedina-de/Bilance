package main

import (
	"genuine/framework"
	_ "genuine/prototype/accounting"
	_ "genuine/prototype/assets"
	_ "genuine/prototype/files"
)

func main() {
	framework.Init()
}
