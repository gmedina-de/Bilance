package register

var Models []any
var Icons []string

func Register[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)
}
