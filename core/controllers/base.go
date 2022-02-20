package controllers

import (
	"genuine/config"
	"genuine/core/template"
	"net/http"
)

type Base struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Name    string

	Data         map[string]any
	TemplateName string
	Template     template.Template
}

func (b *Base) Before(request *http.Request, writer http.ResponseWriter, name string) {

	b.Request = request
	b.Writer = writer
	b.Name = name

	b.Data = make(map[string]any)

	lang := "en-US"
	al := b.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		lang = al[:5]
	}
	b.Data["Lang"] = lang
	if b.Data["Title"] == nil || b.Data["Title"] == "" {
		b.Data["Title"] = config.AppName
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

func (b *Base) ParseForm(model any) {
	b.Request.ParseForm()
	//err := decoder.Decode(model, b.Request.PostForm)
	//if err != nil {
	//	err.Error()
	//	println(err.Error())
	//	// Handle error
	//}
}
