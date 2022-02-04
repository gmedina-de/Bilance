package template

import "strings"

type MenuItem struct {
	Name    string
	Icon    string
	Path    string
	SubMenu []*MenuItem
}

func (i *MenuItem) WithChild(name string, icon string, path string) *MenuItem {
	if i.SubMenu == nil {
		i.SubMenu = []*MenuItem{}
	}
	i.SubMenu = append(i.SubMenu, &MenuItem{name, icon, path, nil})
	return i
}

var navigation []*MenuItem

func AddNavigation(name string, icon string, path string) *MenuItem {
	menuItem := &MenuItem{name, icon, path, nil}
	navigation = append(navigation, menuItem)
	return menuItem
}

func getCurrentNavigation(path string, navigation []*MenuItem) *MenuItem {
	pathParts := strings.Split(path, "/")

	var result *MenuItem
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
