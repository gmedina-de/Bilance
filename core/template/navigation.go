package template

import (
	"sort"
	"strings"
)

var Navigation []*MenuItem

type MenuItem struct {
	Name    string
	Icon    string
	Path    string
	SubMenu []*MenuItem
}

func AddNavigation(name string, icon string) *MenuItem {
	menuItem := &MenuItem{name, icon, "/" + name, nil}
	Navigation = append(Navigation, menuItem)
	sort.Slice(Navigation, func(i, j int) bool {
		return Navigation[i].Path < Navigation[j].Path
	})
	return menuItem
}

func (i *MenuItem) WithChild(name string, icon string) *MenuItem {
	if i.SubMenu == nil {
		i.SubMenu = []*MenuItem{}
	}
	i.SubMenu = append(i.SubMenu, &MenuItem{name, icon, i.Path + "/" + name, nil})
	return i
}

func GetCurrentNavigation(path string, navigation []*MenuItem) *MenuItem {
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

func GetCurrentNavigationIndex(path string, navigation []*MenuItem) int {
	pathParts := strings.Split(path, "/")

	var result int
	max := 0
	for j, n := range navigation {
		nParts := strings.Split(n.Path, "/")
		for i, p := range pathParts {
			if i < len(nParts) && p == nParts[i] && max < i {
				max++
				result = j
			}
		}
	}
	return result
}
