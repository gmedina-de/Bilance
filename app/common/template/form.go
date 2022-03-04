package template

import (
	"genuine/core/models"
	"html/template"
	"reflect"
	"strings"
)

var inputTemplate = template.Must(template.New("").Parse(`
<div class="col-md-6">
	<div class="form-floating">
		<input type="{{.Type}}" class="form-control" name="{{.Name}}" id="{{.Id}}" placeholder="{{.Placeholder}}" value="{{.Value}}" {{.Custom}}>
		<label for="{{.Id}}">{{.Label}}</label>
	</div>
</div>
`))

func inputs(v interface{}) template.HTML {
	tpl, err := inputTemplate.Clone()
	if err != nil {
		return ""
	}
	fields := fields(v)
	var html template.HTML
	for _, field := range fields {
		var sb strings.Builder
		err := tpl.Execute(&sb, field)
		if err != nil {
			return ""
		}
		html = html + template.HTML(sb.String())
	}
	return html
}

type field struct {
	Name        string
	Label       string
	Placeholder string
	Type        string
	Id          string
	Value       any
	Custom      string
}

func fields(v interface{}, names ...string) []field {
	rv := models.RealValueOf(v)
	t := rv.Type()
	ret := make([]field, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		rf := rv.Field(i)
		structField := t.Field(i)

		if !rf.CanInterface() || structField.Name == "Valid" {
			continue
		}

		if structField.Name == models.ID {
			continue
		} else if strings.HasSuffix(structField.Name, models.ID) {

		}

		if structField.Type.Kind() == reflect.Ptr && rf.IsNil() {
			rf = reflect.New(structField.Type.Elem()).Elem()
		}
		if rf.Kind() == reflect.Struct {
			ret = append(ret, fields(rf.Interface(), append(names, structField.Name)...)...)
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
