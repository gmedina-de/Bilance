package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/template"
	"reflect"
	"strings"
)

type Controller = web.ControllerInterface

func PackageName(c Controller) string {
	reflectVal := reflect.ValueOf(c)
	ct := reflect.Indirect(reflectVal).Type()
	controllerName := strings.TrimSuffix(ct.PkgPath(), "/controllers")
	controllerName = controllerName[strings.LastIndex(controllerName, "/")+1:]
	return controllerName
}

type BaseController struct {
	web.Controller
}

func (this *BaseController) Prepare() {
	path := this.Ctx.Request.URL.Path
	currentNavigation1 := template.GetCurrentNavigation(path, template.Navigation)

	c, _ := this.GetControllerAndAction()
	this.ViewPath = PackageName(this) + "/views"
	this.TplName = strings.ToLower(c) + ".gohtml"

	if this.Data["Title"] == nil || this.Data["Title"] == "" {
		this.Data["Title"] = web.BConfig.AppName
	}
	this.Data["Path"] = path
	this.Data["Navigation1"] = template.Navigation
	this.Data["CurrentNavigation1"] = currentNavigation1
	this.Data["CurrentNavigation1Index"] = template.GetCurrentNavigationIndex(path, template.Navigation)
	if currentNavigation1 != nil {
		this.Data["Navigation2"] = currentNavigation1.SubMenu
		this.Data["CurrentNavigation2"] = template.GetCurrentNavigation(path, currentNavigation1.SubMenu)
		this.Data["CurrentNavigation2Index"] = template.GetCurrentNavigationIndex(path, currentNavigation1.SubMenu)
	}

}

func (this *BaseController) Finish() {
}
