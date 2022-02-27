package navigation

import (
	"genuine/app/assets/models"
	"genuine/core/controllers"
	"genuine/core/decorator"
	models2 "genuine/core/models"
	"strconv"
	"strings"
)

type standard struct {
	tree []*Item
}

func Standard() decorator.Decorator {
	s := &standard{}
	s.add("home", "home").Path = "/"
	s.add("accounting", "book").
		Sub("payments", "layers").
		Sub("categories", "tag").
		Sub("analysis", "").
		Sub("balances", "activity").
		Sub("expenses", "").
		Sub("expenses/by_period", "bar-chart-2").
		Sub("expenses/by_category", "pie-chart").
		Path = "/accounting/payments"
	menuItem := s.add("assets", "box")
	Models := models.Models
	for i, m := range Models {
		menuItem = menuItem.Sub(models2.Plural(m), models.Icons[i])
	}
	menuItem.Path = "/assets/" + models2.Plural(Models[0])
	s.add("files", "folder").
		Sub("all", "folder").
		Sub("favorites", "star").
		Sub("last", "clock").
		Sub("trash", "trash")
	s.add("sites", "layout")
	s.add("tasks", "check-circle")
	s.add("settings", "settings").
		Sub("users", "users").
		Path = "/settings/users"
	return s
}

func (s *standard) Decorate(req controllers.Request, res controllers.Response) {
	traverse(res, req.URL.Path, s.tree, 0)
}

func (s *standard) add(name string, icon string) *Item {
	item := &Item{Name: name, Icon: icon, Path: "/" + name}
	s.tree = append(s.tree, item)
	return item
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
