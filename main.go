package main

import (
	_ "homecloud/accounting"
	"homecloud/core"
	_ "homecloud/dashboard"
)

func main() {
	core.Init()
}
