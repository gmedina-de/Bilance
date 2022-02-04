package main

import (
	_ "homecloud/accounting"
	_ "homecloud/assets"
	"homecloud/core"
)

func main() {
	core.Init()
}
