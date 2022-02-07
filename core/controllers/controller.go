package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"homecloud/core/template"
	"strings"
)

type Controller interface {
	web.ControllerInterface
	Routing(Router)
}

type BaseController struct {
	web.Controller
	i18n.Locale
}

func (this *BaseController) Prepare() {
	this.Lang = "en-US"
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5]
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}
	this.Data["Lang"] = this.Lang

	if this.Data["Title"] == nil || this.Data["Title"] == "" {
		this.Data["Title"] = web.BConfig.AppName
	}

	path := this.Ctx.Request.URL.Path
	this.Data["Path"] = path

	currentNavigation1 := template.GetCurrentNavigation(path, template.Navigation)
	this.Data["Navigation1"] = template.Navigation
	this.Data["CurrentNavigation1"] = currentNavigation1
	this.Data["CurrentNavigation1Index"] = template.GetCurrentNavigationIndex(path, template.Navigation)
	if currentNavigation1 != nil {
		this.Data["Navigation2"] = currentNavigation1.SubMenu
		this.Data["CurrentNavigation2"] = template.GetCurrentNavigation(path, currentNavigation1.SubMenu)
		this.Data["CurrentNavigation2Index"] = template.GetCurrentNavigationIndex(path, currentNavigation1.SubMenu)
	}

	c, _ := this.GetControllerAndAction()
	this.TplName = strings.ToLower(c) + ".gohtml"
}

func (this *BaseController) Finish() {
}
