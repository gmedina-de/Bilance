package functions

import (
	"genuine/models"
	"html"
	"html/template"
	"strings"
)

type table struct {
	translator Translator
}

func Table(translator Translator) Provider {
	return &table{translator}
}

func (t *table) GetFuncMap() template.FuncMap {
	return map[string]any{
		"th": t.th,
		"td": t.td,
	}
}

func (t *table) th(v any) []string {
	rv := models.RealValueOf(v)
	rt := rv.Type()
	ret := make([]string, 0, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		sf := rt.Field(i)
		if !sf.IsExported() || strings.HasSuffix(sf.Name, models.ID) && sf.Name != models.ID {
			continue
		}
		if sf.Name == "Model" {
			ret = append(ret, models.ID)
			continue
		}
		ret = append(ret, t.translator.Translate(sf.Name))
	}
	return ret
}

func (t *table) td(v any) template.HTML {
	tpl, _ := tdTemplate.Clone()
	rv := models.RealValueOf(v)
	rt := rv.Type()
	var ret template.HTML
	for i := 0; i < rt.NumField(); i++ {
		rf := rv.Field(i)
		sf := rt.Field(i)
		if !sf.IsExported() || strings.HasSuffix(sf.Name, models.ID) && sf.Name != models.ID {
			continue
		}
		var sb strings.Builder
		if sf.Name == "Model" {
			tpl.Execute(&sb, rf.FieldByName(models.ID).Uint())
		} else {
			tpl.Execute(&sb, rf.Interface())
		}

		ret = ret + template.HTML(html.UnescapeString(sb.String()))
	}

	return ret
}

var tdTemplate = template.Must(template.New("").Parse("<td>{{.}}</td>"))
