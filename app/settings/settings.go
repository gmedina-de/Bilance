package settings

import (
	"genuine/app/settings/authenticator"
	"genuine/app/settings/controllers"
	"genuine/app/settings/repositories"
	"genuine/core"
	"genuine/core/models"
	template2 "genuine/core/template"
	"html/template"
	"reflect"
	"strings"
)

func init() {
	core.Provide(controllers.Users)
	core.Provide(repositories.Users)
	core.Provide(authenticator.Basic)

	template2.AddFunc("td", td)
	template2.AddFunc("th", th)
	template2.AddFunc("paginate", paginate)
	template2.AddFunc("inputs", inputs)
}

var tdTemplate = template.Must(template.New("").Parse("<td>{{.}}</td>"))

func td(v interface{}) (template.HTML, error) {
	tpl, err := tdTemplate.Clone()
	if err != nil {
		return "", err
	}
	var html template.HTML
	rv := models.RealValueOf(v)
	t := rv.Type()
	for i := 0; i < t.NumField(); i++ {
		rf := rv.Field(i)
		structField := t.Field(i)
		if structField.Type.Kind() == reflect.Ptr && rf.IsNil() {
			rf = reflect.New(structField.Type.Elem()).Elem()
		}
		if rf.Kind() == reflect.Struct {
			var t2 template.HTML
			t2, _ = td(rf.Interface())
			html = html + t2
			continue
		}
		var sb strings.Builder
		err := tpl.Execute(&sb, rf.Interface())
		if err != nil {
			return "", err
		}
		html = html + template.HTML(sb.String())
	}

	if err != nil {
		return "", err
	}
	return html, nil
}

func th(v interface{}) []string {
	rv := models.RealValueOf(v)
	t := rv.Type()
	ret := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		rf := rv.Field(i)
		structField := t.Field(i)
		if structField.Type.Kind() == reflect.Ptr && rf.IsNil() {
			rf = reflect.New(structField.Type.Elem()).Elem()
		}
		if rf.Kind() == reflect.Struct {
			ret = append(ret, th(rf.Interface())...)
			continue
		}
		ret = append(ret, structField.Name)
	}
	return ret
}

func paginate(pages int64, page int64, offset int64) []int64 {
	var i int64
	var items []int64
	for i = page - offset; i <= page+offset; i++ {
		if i <= pages && i > 0 {
			items = append(items, i)
		}
	}
	return items
}

var inputTemplate = template.Must(template.New("").Parse(`
<div class="col-md-6">
	<div class="form-floating">
		<input type="{{.Type}}" class="form-control" name="{{.Name}}" id="{{.Id}}" placeholder="{{.Placeholder}}" value="{{.Value}}" {{.Custom}}>
		<label for="{{.Id}}">{{.Label}}</label>
	</div>
</div>
`))

func inputs(v interface{}, errs ...error) (template.HTML, error) {
	tpl, err := inputTemplate.Clone()
	if err != nil {
		return "", err
	}
	fields := Fields(v)
	var html template.HTML
	for _, field := range fields {
		var sb strings.Builder
		err := tpl.Execute(&sb, field)
		if err != nil {
			return "", err
		}
		html = html + template.HTML(sb.String())
	}
	return html, nil
}

type field struct {
	Name        string
	Label       string
	Placeholder string
	Type        string
	Id          string
	Value       interface{}
	Custom      string
}

func Fields(v interface{}, names ...string) []field {
	rv := models.RealValueOf(v)
	t := rv.Type()
	ret := make([]field, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		rf := rv.Field(i)
		structField := t.Field(i)

		if structField.Name == "Id" {
			continue
		}
		if structField.Type.Kind() == reflect.Ptr && rf.IsNil() {
			rf = reflect.New(structField.Type.Elem()).Elem()
		}
		if rf.Kind() == reflect.Struct {
			ret = append(ret, Fields(rf.Interface(), append(names, structField.Name)...)...)
			continue
		}
		name := append(names, structField.Name)

		finalName := strings.Join(name, ".")
		f := field{
			Name:        finalName,
			Label:       structField.Name,
			Placeholder: structField.Name,
			Type:        fieldInputType(structField),
			Id:          finalName,
			Value:       rf.Interface(),
			Custom:      structField.Tag.Get("form"),
		}
		ret = append(ret, f)
	}
	return ret
}

func fieldInputType(t reflect.StructField) string {
	name := t.Name
	kind := t.Type.Kind()

	switch name {
	case "Color":
		return "color"
	case "Date":
		return "date"
	case "Email":
		return "email"
	case "File":
		return "file"
	case "Image":
		return "image"
	case "Month":
		return "month"
	case "Number":
		return "number"
	case "Password":
		return "password"
	case "Range":
		return "range"
	case "Search":
		return "search"
	case "Tel":
		return "tel"
	case "Time":
		return "time"
	case "Url":
		return "url"
	case "Week":
		return "week"
	}
	if kind == reflect.Bool {
		return "checkbox"
	}
	if kind == reflect.Int || kind == reflect.Int64 {
		return "number"
	}
	return "text"
}
