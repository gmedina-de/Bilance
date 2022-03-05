package template

import (
	"genuine/core/template"
)

func init() {
	template.AddFunc("td", td)
	template.AddFunc("th", th)
	template.AddFunc("paginate", paginate)
	template.AddFunc("inputs", inputs)
}
