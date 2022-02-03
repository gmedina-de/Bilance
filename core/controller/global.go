package controller

import (
	"homecloud/core/localization"
	"homecloud/core/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Context struct {
	User       *model.User
	Path       string
	Title      string
	Navigation []MenuItem
	Parameters *Parameters
}

type Parameters struct {
	Model      interface{}
	Data       interface{}
	Pagination *Pagination
	Toast      string
}

type Pagination struct {
	Pages  int64
	Active int64
}

type MenuItem struct {
	Name string
	Icon string
	Path string
}

var menu []MenuItem

func AddMenuItem(name string, icon string, path string) {
	menu = append(menu, MenuItem{name, icon, path})
}

func Render(writer http.ResponseWriter, request *http.Request, parameters *Parameters, title string, templates ...string) {
	// prepare templates
	templates = append(templates,
		"core/template/base.gohtml",
		"core/template/navbar.gohtml",
		"core/template/navigation.gohtml",
	)
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"translate": localization.Translate,
		"active":    active,
		"paginate":  paginate,
		"sum":       sum,
		"contains":  contains,
	})
	tmpl, err := tmpl.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}

	var user model.User
	model.Deserialize(request.Header.Get("user"), &user)
	err = tmpl.ExecuteTemplate(writer, "base", &Context{
		&user,
		request.URL.Path,
		title,
		menu,
		parameters,
	})
	if err != nil {
		panic(err)
	}
}

func Redirect(writer http.ResponseWriter, request *http.Request, path string) {
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

func contains(a string, b int64) bool {
	return strings.Contains(a, strconv.FormatInt(b, 10))
}
