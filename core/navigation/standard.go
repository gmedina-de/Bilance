package navigation

import (
	"genuine/core/http"
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

func (s *standard) Handle(response http.Response) {
	path := response["Path"].(string)

	currentNavigation1 := s.getCurrentNavigation(path, s.tree)
	response["Navigation1"] = s.tree
	response["CurrentNavigation1"] = currentNavigation1
	response["CurrentNavigation1Index"] = s.getCurrentNavigationIndex(path, s.tree)
	if currentNavigation1 != nil {
		response["Navigation2"] = currentNavigation1.SubMenu
		response["CurrentNavigation2"] = s.getCurrentNavigation(path, currentNavigation1.SubMenu)
		response["CurrentNavigation2Index"] = s.getCurrentNavigationIndex(path, currentNavigation1.SubMenu)
	}
}

func (s *standard) getCurrentNavigation(path string, navigation []*Item) *Item {
	pathParts := strings.Split(path, "/")

	var result *Item
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

func (s *standard) getCurrentNavigationIndex(path string, navigation []*Item) int {
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
