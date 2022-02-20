package controllers

import (
	"genuine/core/template"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"github.com/gorilla/schema"
	"net/http"
)

type Base struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Name    string

	Data         map[string]any
	TemplateName string
	Template     template.Template

	i18n.Locale
}

func (this *Base) Before(request *http.Request, writer http.ResponseWriter, name string) {

	this.Request = request
	this.Writer = writer
	this.Name = name

	this.Data = make(map[string]any)

	this.Lang = "en-US"
	al := this.Request.Header.Get("Accept-Language")
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

	path := this.Request.URL.Path
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

	this.TemplateName = this.Name + ".gohtml"

}

func (this *Base) After() {
	this.Template.Render(this.Writer, this.TemplateName, this.Data)
}

func (this *Base) Redirect(url string, status int) {
	http.Redirect(this.Writer, this.Request, url, status)
}

var decoder = schema.NewDecoder()

func (this *Base) ParseForm(model any) {
	this.Request.ParseForm()
	err := decoder.Decode(model, this.Request.PostForm)
	if err != nil {
		err.Error()
		println(err.Error())
		// Handle error
	}
}
