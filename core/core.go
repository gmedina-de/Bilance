package core

import (
	"genuine/core/authenticator"
	"genuine/core/controllers"
	"genuine/core/database"
	"genuine/core/injector"
	"genuine/core/log"
	"genuine/core/models"
	"genuine/core/repositories"
	"genuine/core/template"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"strconv"
	"strings"
)

func init() {

	template.AddNavigation("home", "home").
		Path = "/"

	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/settings/users"

	injector.Implementations(
		log.Console,
		database.Orm,
		repositories.Users,
		controllers.Index,
		controllers.Users,
		authenticator.Basic,
	)

	orm.RegisterModel(
		&models.User{},
	)

	web.BConfig.AppName = "HomeCloud"
	web.BConfig.Listen.HTTPAddr = "0.0.0.0"
	web.BConfig.Listen.HTTPPort = 8080
	web.BConfig.RunMode = web.DEV
	web.BConfig.WebConfig.AutoRender = true
	web.BConfig.RecoverPanic = false
	web.BConfig.Listen.EnableAdmin = false

	web.AddFuncMap("td", template.Td)
	web.AddFuncMap("th", template.Th)
	web.AddFuncMap("i18n", i18n.Tr)
	web.AddFuncMap("sum", func(a int, b int) int { return a + b })
	web.AddFuncMap("contains", func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) })

	web.ExceptMethodAppend("Routing")

	languages := []string{"en-US", "de-DE"}
	for _, lang := range languages {
		if err := i18n.SetMessage(lang, "i18n/"+"locale_"+lang+".ini"); err != nil {
			return
		}
	}
}

func Init() {
	injector.Injector(func(cs []controllers.Controller, auth authenticator.Authenticator, other []any) {
		err := orm.RunSyncdb(database.Name, false, false)
		if err != nil {
			panic(err)
		}

		for _, c := range cs {
			c.Routing()
			web.Include(c)
		}

		web.Run()
	})
}
