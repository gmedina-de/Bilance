package core

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/application"
	"homecloud/core/authenticator"
	"homecloud/core/controllers"
	"homecloud/core/database"
	"homecloud/core/injector"
	"homecloud/core/log"
	"homecloud/core/repository"
	"homecloud/core/server"
	"homecloud/core/template"
)

func init() {
	web.AddViewPath("core/views")

	template.AddNavigation("home", "home").Path = "/"
	Implementations(
		log.Console,
		database.Gorm,
		authenticator.Basic,
		server.Authenticated,
		repository.Users,
		controllers.Index,
		controllers.Users,
		controllers.Index,
	)
	//web.AddTemplateExt("gohtml") not needed
	err := web.BuildTemplate("core/views")
	if err != nil {
		fmt.Println(err.Error())
	}

	web.BConfig.AppName = "HomeCloud"
	web.BConfig.Listen.HTTPAddr = "0.0.0.0"
	web.BConfig.Listen.HTTPPort = 8080
	web.BConfig.RunMode = "dev"
	web.BConfig.WebConfig.AutoRender = true
	web.BConfig.RecoverPanic = false

}

var inj injector.Injector = injector.Recursive()

func Implementations(constructors ...interface{}) {
	for _, constructor := range constructors {
		inj.Add(constructor)
	}
}

func Init() {
	inj.Inject(application.Homecloud).Interface().(application.Application).Run()
}
