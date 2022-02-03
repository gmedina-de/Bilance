package template

import (
	"homecloud/core/localization"
	"homecloud/core/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Context struct {
	User        *model.User
	Path        string
	Title       string
	Navigation  []*menuItem
	Navigation2 []*menuItem
	Parameters  *Parameters
}

type Parameters struct {
	Model      interface{}
	Data       interface{}
	Pagination *Pagination
	Toast      string
}

func Render(writer http.ResponseWriter, request *http.Request, parameters *Parameters, title string, templates ...string) {
	// prepare templates
	templates = append(templates,
		"core/template/base.gohtml",
		"core/template/navigation1.gohtml",
		"core/template/navigation2.gohtml",
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
		navigation,
		getCurrentNavigation(request.URL.Path).subMenu,
		parameters,
	})
	if err != nil {
		panic(err)
	}
}

func active(currentPath string, linkPath string) bool {
	if currentPath == "/" && linkPath == "/" || strings.HasPrefix(currentPath, linkPath) && linkPath != "/" {
		return true
	}
	return false
}

func sum(a int64, b int64) int64 {
	return a + b
}

func contains(a string, b int64) bool {
	return strings.Contains(a, strconv.FormatInt(b, 10))
}
