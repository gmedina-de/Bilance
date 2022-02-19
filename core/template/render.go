package template

import (
	"github.com/beego/i18n"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
)

var templates = []string{
	"views/base.gohtml",
	"views/navigation1.gohtml",
	"views/navigation2.gohtml",
	"views/pagination.gohtml",
}

var tmpl = template.New(path.Base(templates[0]))

func init() {
	tmpl.Funcs(template.FuncMap{
		"i18n":     i18n.Tr,
		"td":       Td,
		"th":       Th,
		"sum":      func(a int, b int) int { return a + b },
		"contains": func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) },
	})
}

func Render(writer http.ResponseWriter, template string, data map[string]any) {
	tmpl.ParseFiles(append(templates, template)...)
	err := tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
}
