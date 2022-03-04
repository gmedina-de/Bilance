package template

import (
	"genuine/core/models"
	"html/template"
	"strings"
)

func th(v any) []string {
	rv := models.RealValueOf(v)
	t := rv.Type()
	ret := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		rf := rv.Field(i)
		field := t.Field(i)

		if !rf.CanInterface() ||
			field.Name == "Valid" ||
			strings.HasSuffix(field.Name, models.ID) && field.Name != models.ID {
			continue
		}

		if field.Name == "Model" {
			ret = append(ret, models.ID)
			continue
		}

		ret = append(ret, field.Name)

	}
	return ret
}

func td(v any) template.HTML {
	tpl, _ := tdTemplate.Clone()
	rv := models.RealValueOf(v)
	t := rv.Type()
	var ret template.HTML
	for i := 0; i < t.NumField(); i++ {
		rf := rv.Field(i)
		field := t.Field(i)

		if !rf.CanInterface() ||
			field.Name == "Valid" ||
			strings.HasSuffix(field.Name, models.ID) && field.Name != models.ID {
			continue
		}

		var sb strings.Builder
		if field.Name == "Model" {
			tpl.Execute(&sb, rf.FieldByName(models.ID).Uint())
		} else {
			tpl.Execute(&sb, rf.Interface())
		}
		ret = ret + template.HTML(sb.String())
	}

	return ret
}

var tdTemplate = template.Must(template.New("").Parse("<td>{{.}}</td>"))
