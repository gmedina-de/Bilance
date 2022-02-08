package core

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"homecloud/core/controllers"
	"homecloud/core/database"
	"homecloud/core/filters"
	"homecloud/core/log"
	"homecloud/core/repositories"
	"homecloud/core/template"
	"strconv"
	"strings"
)

func init() {

	template.AddNavigation("home", "home").
		Path = "/"

	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/settings/users"

	Implementations(
		log.Console,
		database.Orm,
		repositories.Users,
		controllers.Index,
		controllers.Users,
		filters.Auth,
	)

	web.BConfig.AppName = "HomeCloud"
	web.BConfig.Listen.HTTPAddr = "0.0.0.0"
	web.BConfig.Listen.HTTPPort = 8080
	web.BConfig.RunMode = web.DEV
	web.BConfig.WebConfig.AutoRender = true
	web.BConfig.RecoverPanic = false
	web.BConfig.Listen.EnableAdmin = true

	web.AddFuncMap("td", template.Td)
	web.AddFuncMap("th", template.Th)
	web.AddFuncMap("i18n", i18n.Tr)
	web.AddFuncMap("sum", func(a int, b int) int { return a + b })
	web.AddFuncMap("contains", func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) })

	web.AddViewPath("core/views")
	web.ExceptMethodAppend("Routing")

	languages := []string{"en-US", "de-DE"}
	for _, lang := range languages {
		if err := i18n.SetMessage(lang, "i18n/"+"locale_"+lang+".ini"); err != nil {
			return
		}
	}
}

func Init() {
	Injector(func(cs []controllers.Controller, fs []filters.Filter) {
		for _, c := range cs {
			c.Routing()
			web.Include(c)
		}
		for _, f := range fs {
			web.InsertFilter(f.Pattern(), f.Pos(), f.Func(), f.Opts()...)
		}
		web.Run()
	})
}
