package navigation

type item struct {
	Name    string
	Icon    string
	Path    string
	SubMenu func() items
}

type items []*item

func (is *items) add(name string, icon string, path string) *item {
	i := &item{Name: name, Icon: icon, Path: path}
	*is = append(*is, i)
	return i
}
