package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/server"
	"homecloud/core/template"
)

type ControllerOld interface {
	Routing(server server.Server)
}
type Controller interface {
	web.ControllerInterface
	Route() string
}

type BaseController struct {
	web.Controller
}

func (this *BaseController) Prepare() {
	this.Data["Title"] = "Hello"

	path := this.Ctx.Request.URL.Path
	currentNavigation1 := template.GetCurrentNavigation(path, template.Navigation)
	this.Data["Navigation1"] = template.Navigation
	this.Data["Navigation2"] = currentNavigation1.SubMenu
	this.Data["CurrentNavigation1"] = currentNavigation1
	this.Data["CurrentNavigation2"] = template.GetCurrentNavigation(path, currentNavigation1.SubMenu)
	this.Data["CurrentNavigation1Index"] = template.GetCurrentNavigationIndex(path, template.Navigation)
	this.Data["CurrentNavigation2Index"] = template.GetCurrentNavigationIndex(path, currentNavigation1.SubMenu)

}

func (this *BaseController) Finish() {
	this.Data["Title"] = "Hello"
}
