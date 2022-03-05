package navigation

import (
	model2 "genuine/app/accounting/models"
	"genuine/app/assets/models"
	"genuine/core/controllers"
	"genuine/core/decorator"
	models2 "genuine/core/models"
	"genuine/core/repositories"
	"strconv"
	"strings"
)

type standard struct {
	root func() items
}

func Standard(categories repositories.Repository[model2.Category]) decorator.Decorator {
	return &standard{root: func() items {
		var items0 items
		items0.add("home", "home", "/")

		items0.add("accounting", "book", "/accounting/payments").SubMenu = func() items {
			var items1 items
			items1.add("payments", "layers", "/accounting/payments")
			items1.add("categories", "tag", "/accounting/categories").SubMenu = func() items {
				var items2 items
				for _, c := range categories.All() {
					items2.add(c.Name, "search", "/accounting/payments?q=category_id:"+strconv.FormatUint(uint64(c.ID), 10))
				}
				return items2
			}
			items1.add("analysis", "", "analysis")
			items1.add("balances", "activity", "balances")
			items1.add("expenses", "", "expenses")
			items1.add("expenses/by_period", "bar-chart-2", "expenses/by_period")
			items1.add("expenses/by_category", "pie-chart", "expenses/by_category")
			return items1
		}

		Models := models.Models
		items0.add("assets", "box", "/assets"+models2.Plural(Models[0])).SubMenu = func() items {
			var items1 items
			for i, m := range Models {
				items1.add(models2.Plural(m), models.Icons[i], "/assets/"+models2.Plural(m))
			}
			return items1
		}

		items0.add("files", "folder", "/files").SubMenu = func() items {
			var items1 items
			items1.add("all", "folder", "/files/all")
			items1.add("favorites", "star", "/files/favorites")
			items1.add("last", "clock", "/files/last")
			items1.add("trash", "trash", "/files/trash")
			return items1
		}

		items0.add("sites", "layout", "/sites")

		items0.add("tasks", "check-circle", "/tasks")

		items0.add("settings", "settings", "/settings/users").SubMenu = func() items {
			var items1 items
			items1.add("users", "users", "/settings/users")
			return items1
		}
		return items0
	}}
}

func (s *standard) Decorate(req controllers.Request, res controllers.Response) {
	traverse(res, req.URL.Path, s.root(), 0)
}

func traverse(response controllers.Response, path string, tree items, level int) {
	currentNavigation := currentNavigation(path, tree)
	levelString := strconv.Itoa(level)
	response["Navigation"+levelString] = tree
	response["CurrentNavigation"+levelString] = currentNavigation
	if currentNavigation != nil && currentNavigation.SubMenu != nil {
		level++
		traverse(response, path, currentNavigation.SubMenu(), level)
	}
}

func currentNavigation(path string, navigation items) *item {
	pathParts := strings.Split(path, "/")
	var result *item
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
