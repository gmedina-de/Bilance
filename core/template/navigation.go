package template

import "strings"

type menuItem struct {
	Name    string
	Icon    string
	Path    string
	subMenu []*menuItem
}

func (i *menuItem) WithSubItems(menuItems ...*menuItem) *menuItem {
	i.subMenu = menuItems
	return i
}

func MenuItem(name string, icon string, path string) *menuItem {
	return &menuItem{Name: name, Icon: icon, Path: path}
}

var navigation []*menuItem

func AddNavigation(item *menuItem) {
	navigation = append(navigation, item)
}

func getCurrentNavigation(path string) *menuItem {
	for _, n := range navigation {
		if strings.HasPrefix(path, n.Path) && n.subMenu != nil {
			return n
		}
	}
	return navigation[0]
}
