package template

import (
	"genuine/core/models"
	"html/template"
	"reflect"
	"strings"
)

var tdTemplate = template.Must(template.New("").Parse("<td>{{.}}</td>"))

func Td(v interface{}) (template.HTML, error) {
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
			t2, _ = Td(rf.Interface())
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

func Th(v interface{}) []string {
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
			ret = append(ret, Th(rf.Interface())...)
			continue
		}
		ret = append(ret, structField.Name)
	}
	return ret
}
