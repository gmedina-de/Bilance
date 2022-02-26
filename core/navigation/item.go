package navigation

type Item struct {
	Name    string
	Icon    string
	Path    string
	SubMenu []*Item
}

func (i *Item) Sub(name string, icon string) *Item {
	if i.SubMenu == nil {
		i.SubMenu = []*Item{}
	}
	i.SubMenu = append(i.SubMenu, &Item{Name: name, Icon: icon, Path: "/" + name})
	return i
}
