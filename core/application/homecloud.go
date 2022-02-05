package application

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/controllers"
	"homecloud/core/server"
	"homecloud/core/template"
)

type homecloud struct {
	controllers  []controllers.ControllerOld
	controllers2 []controllers.Controller
	server       server.Server
}

func Homecloud(server server.Server, controllers []controllers.ControllerOld, controllers2 []controllers.Controller) *homecloud {
	return &homecloud{controllers, controllers2, server}
}

func (b *homecloud) Run() {

	// adds settings to the end of navigation
	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/settings/users"

	for _, c := range b.controllers2 {
		web.Router(c.Route(), c)
		//c.Routing(b.server)
	}
	web.Run()

	//b.server.Start()
}
