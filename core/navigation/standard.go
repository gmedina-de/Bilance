package navigation

import (
	"genuine/core/controllers"
	"strconv"
	"strings"
)

type standard struct {
	tree []*Item
}

func Standard() Navigation {
	return &standard{}
}

func (s *standard) Add(name string, icon string) *Item {
	item := &Item{Name: name, Icon: icon, Path: "/" + name}
	s.tree = append(s.tree, item)
	return item
}

func (s *standard) Handle(response controllers.Response) {
	path := response["Path"].(string)
	traverse(response, path, s.tree, 0)
}

func traverse(response controllers.Response, path string, tree []*Item, level int) {
	currentNavigation := currentNavigation(path, tree)
	levelString := strconv.Itoa(level)
	response["Navigation"+levelString] = tree
	response["CurrentNavigation"+levelString] = currentNavigation
	if currentNavigation != nil && currentNavigation.SubMenu != nil {
		level++
		traverse(response, path, currentNavigation.SubMenu, level)
	}
}

func currentNavigation(path string, navigation []*Item) *Item {
	pathParts := strings.Split(path, "/")

	var result *Item
	max := 0
	for _, n := range navigation {
		nParts := strings.Split(n.Path, "/")
		for j, p := range pathParts {
			if j < len(nParts) && p == nParts[j] && max < j {
				max++
				result = n
			}
		}
	}
	return result
}
