package core

import (
	"github.com/beego/beego/v2/server/web"
	. "homecloud/core/controllers"
	"homecloud/core/database"
	"homecloud/core/injector"
	log2 "homecloud/core/log"
	"homecloud/core/repositories"
	"homecloud/core/template"
)

func init() {
	template.AddNavigation("home", "home").
		Path = "/"

	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/settings/users"

	_ = web.AddViewPath("core/views")

	injector.Implementations(
		log2.Console,
		database.Orm,
		repositories.Users,
		Index,
		Users,
	)

	web.BConfig.AppName = "HomeCloud"
	web.BConfig.Listen.HTTPAddr = "0.0.0.0"
	web.BConfig.Listen.HTTPPort = 8080
	web.BConfig.RunMode = web.DEV
	web.BConfig.WebConfig.AutoRender = true
	web.BConfig.RecoverPanic = false
}

func Init() {
	injector.Inj.Inject(homecloud)
}

func homecloud(controllers []Controller) {
	for _, c := range controllers {
		c.Routing()
	}
	web.Run()
}
