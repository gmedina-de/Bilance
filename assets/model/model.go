package model

var Models []any
var Icons []string

func AddModel(model any, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)
}
