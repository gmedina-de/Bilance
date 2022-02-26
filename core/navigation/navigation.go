package navigation

type Navigation interface {
	Add(name string, icon string) *Item
	Handle(data map[string]any)
}
