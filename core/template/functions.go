package template

import "html/template"

var funcMap = template.FuncMap{}

func AddFunc(name string, function any) {
	funcMap[name] = function
}

func FuncMap() template.FuncMap {
	return funcMap
}
