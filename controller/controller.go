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
	User              *model.User
	SelectedProjectId int64
	Path              string
	Title             string
	Parameters        *Parameters
}

type Parameters struct {
	Model interface{}
	Data  interface{}
	Toast string
}

func render(writer http.ResponseWriter, request *http.Request, parameters *Parameters, title string, templates ...string) {
	// prepare templates
	for i, _template := range templates {
		templates[i] = "template/" + _template + ".gohtml"
	}
	templates = append(templates,
		"template/base.gohtml",
		"template/navbar.gohtml",
		"template/navigation.gohtml",
	)
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"translate": localization.Translate,
		"active":    active,
		"paginate":  paginate,
		"sum":       sum,
	})
	tmpl, err := tmpl.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	user := model.DeserializeUser(request.Header.Get("user"))
	err = tmpl.ExecuteTemplate(writer, "base", &Context{
		user,
		model.GetSelectedProjectId(request),
		request.URL.Path,
		title,
		parameters,
	})
	if err != nil {
		panic(err)
	}
}

func redirect(writer http.ResponseWriter, request *http.Request, path string) {
	http.Redirect(writer, request, path, http.StatusTemporaryRedirect)
}

func active(currentPath string, linkPath string) string {
	if currentPath == "/" && linkPath == "/" || strings.HasPrefix(currentPath, linkPath) && linkPath != "/" {
		return " active"
	}
	return ""
}

func paginate(count int64) []int64 {
	var i int64
	var items []int64
	for i = 1; i <= count; i++ {
		items = append(items, i)
	}
	return items
}

func sum(a int64, b int64) int64 {
	return a + b
}
