package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/service"
	"html/template"
	"net/http"
	"strings"
)

type Controller interface {
	Routing(router service.Router)
}

type Context struct {
	Data  interface{}
	Title string
	User  *model.User
	Path  string
}

func render(writer http.ResponseWriter, request *http.Request, title string, data interface{}, templates ...string) {
	for i, iTemplate := range templates {
		templates[i] = "view/" + iTemplate + ".gohtml"
	}
	templates = append(templates, "view/base.gohtml", "view/navbar.gohtml", "view/navigation.gohtml")
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"translate": localization.Translate,
		"active":    Active,
	})
	tmpl, err := tmpl.ParseFiles(templates...)
	context := &Context{
		data,
		title,
		model.DeserializeUser(request.Header.Get("user")),
		request.URL.Path,
	}
	err = tmpl.ExecuteTemplate(writer, "base", context)
	if err != nil {
		panic(err)
	}
}

func Active(currentPath string, linkPath string) string {
	if currentPath == "/" && linkPath == "/" || strings.HasPrefix(currentPath, linkPath) && linkPath != "/" {
		return " active"
	}
	return ""
}
