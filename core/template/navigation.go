package template

import "strings"

type menuItem struct {
	Name    string
	Icon    string
	Path    string
	SubMenu []*menuItem
}

func MenuItem(name string, icon string, path string) *menuItem {
	return &menuItem{Name: name, Icon: icon, Path: path}
}

func (i *menuItem) WithSubItems(menuItems ...*menuItem) *menuItem {
	i.SubMenu = menuItems
	return i
}

var navigation []*menuItem

func AddNavigation(item *menuItem) {
	navigation = append(navigation, item)
}

func getCurrentNavigation(path string, navigation []*menuItem) *menuItem {
	pathParts := strings.Split(path, "/")

	var result *menuItem
	max := 0

	for _, n := range navigation {
		nParts := strings.Split(n.Path, "/")
		for i, p := range pathParts {
			if i < len(nParts) && p == nParts[i] && max < i {
				max++
				result = n
			}
		}
	}
	return result
}
