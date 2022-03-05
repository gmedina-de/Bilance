package functions

import "html/template"

type Provider interface {
	GetFuncMap() template.FuncMap
}
