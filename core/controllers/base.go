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

func (b *Base) Before(request *http.Request, writer http.ResponseWriter, name string) {

	b.Request = request
	b.Writer = writer
	b.Name = name

	b.Data = make(map[string]any)

	b.Lang = "en-US"
	al := b.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5]
		if i18n.IsExist(al) {
			b.Lang = al
		}
	}
	b.Data["Lang"] = b.Lang

	if b.Data["Title"] == nil || b.Data["Title"] == "" {
		b.Data["Title"] = web.BConfig.AppName
	}

	path := b.Request.URL.Path
	b.Data["Path"] = path

	currentNavigation1 := template.GetCurrentNavigation(path, template.Navigation)
	b.Data["Navigation1"] = template.Navigation
	b.Data["CurrentNavigation1"] = currentNavigation1
	b.Data["CurrentNavigation1Index"] = template.GetCurrentNavigationIndex(path, template.Navigation)
	if currentNavigation1 != nil {
		b.Data["Navigation2"] = currentNavigation1.SubMenu
		b.Data["CurrentNavigation2"] = template.GetCurrentNavigation(path, currentNavigation1.SubMenu)
		b.Data["CurrentNavigation2Index"] = template.GetCurrentNavigationIndex(path, currentNavigation1.SubMenu)
	}

	b.TemplateName = b.Name + ".gohtml"

}

func (b *Base) After() {
	b.Template.Render(b.Writer, b.TemplateName, b.Data)
}

func (b *Base) Redirect(url string, status int) {
	http.Redirect(b.Writer, b.Request, url, status)
}

var decoder = schema.NewDecoder()

func (b *Base) ParseForm(model any) {
	b.Request.ParseForm()
	err := decoder.Decode(model, b.Request.PostForm)
	if err != nil {
		err.Error()
		println(err.Error())
		// Handle error
	}
}
