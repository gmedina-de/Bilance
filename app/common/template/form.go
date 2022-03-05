package template

import (
	"fmt"
	"genuine/core/database"
	"genuine/core/models"
	"html/template"
	"reflect"
	"strings"
)

var templates = map[string]*template.Template{
	"default": parse(`
<div class="form-floating mb-3">
	<input type="{{.Type}}" class="form-control" name="{{.Name}}" id="{{.Id}}" placeholder="{{.Placeholder}}" value="{{.Value}}" {{.Custom}}>
	<label for="{{.Id}}">{{.Label}}</label>
</div>
	`),
	"select": parse(`
<div class="form-floating mb-3">
	<select class="form-select" name="{{.Name}}ID" id="{{.Id}}" {{.Custom}}>
		<option></option>
		{{range .Options}}
			<option value="{{.Value}}"{{if .Selected}} selected{{end}}>{{.Label}}</option>
		{{end}}
	</select>
	<label for="{{.Id}}">{{.Label}}</label>
</div>
	`),
	"checkbox": parse(`
<div class="form-check mb-3">
  <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault">
  <label class="form-check-label" for="flexCheckDefault">Default checkbox</label>
</div>
	`),
}

func parse(parse string) *template.Template {
	return template.Must(template.New("").Parse(parse))
}

type field struct {
	Name        string
	Label       string
	Placeholder string
	Id          string
	Type        string
	Value       any
	Options     []option
	Custom      string
}

type option struct {
	Value    uint
	Selected bool
	Label    string
}

func inputs(model any, database database.Database) template.HTML {
	fields := fields(model, database)
	var html template.HTML
	for _, field := range fields {
		var sb strings.Builder
		tmpl, found := templates[field.Type]
		if !found {
			tmpl = templates["default"]
		}
		tmpl.Execute(&sb, field)
		html = html + template.HTML(sb.String())
	}
	return html
}

func fields(model any, database database.Database) []field {
	rv := models.RealValueOf(model)
	t := rv.Type()
	ret := make([]field, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		rf := rv.Field(i)
		sf := t.Field(i)
		if !sf.IsExported() || sf.Name == "Model" || strings.HasSuffix(sf.Name, models.ID) {
			continue
		}
		ret = append(ret, field{
			Name:        sf.Name,
			Label:       sf.Name,
			Placeholder: sf.Name,
			Id:          sf.Name,
			Type:        fieldInputType(sf),
			Value:       rf.Interface(),
			Options:     options(sf, selectedId(rv, sf), database),
			Custom:      sf.Tag.Get("form"),
		})
	}
	return ret
}

func selectedId(rv reflect.Value, sf reflect.StructField) uint {
	idField := rv.FieldByName(sf.Name + models.ID)
	if !reflect.ValueOf(idField).IsZero() {
		return uint(idField.Uint())
	}
	return 0
}

func options(sf reflect.StructField, selectedId uint, database database.Database) []option {
	var ret []option
	slice := reflect.MakeSlice(reflect.SliceOf(sf.Type), 0, 10).Interface()
	database.Find(&slice)
	sv := reflect.ValueOf(slice)
	for i := 0; i < sv.Len(); i++ {
		item := sv.Index(i)
		id := uint(item.FieldByName(models.ID).Uint())
		ret = append(ret, option{
			Value:    id,
			Selected: id == selectedId,
			Label:    fmt.Sprintf("%s", item.Interface()),
		})
	}
	return ret
}

func fieldInputType(t reflect.StructField) string {
	switch t.Name {
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
	switch t.Type.Kind() {
	case reflect.Bool:
		return "checkbox"
	case reflect.Int | reflect.Int64:
		return "number"
	case reflect.Struct:
		return "select"
	}
	return "text"
}
