package main

import (
	_ "homecloud/accounting"
	_ "homecloud/assets"
	"homecloud/core"
	_ "homecloud/files"
)

func main() {
	core.Init()
}
