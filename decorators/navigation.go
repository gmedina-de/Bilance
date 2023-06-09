package decorators

import (
	"genuine/controllers"
	models2 "genuine/models"
	"genuine/repositories"
	"strconv"
	"strings"
)

type navigation struct {
	root func() items
}

func Navigation(
	categories repositories.Repository[models2.Category],
	sites repositories.Repository[models2.Site],
	assets []models2.Asset,
) Decorator {
	return &navigation{root: func() items {
		var items0 items
		items0.add("home", "home", "/")

		items0.add("accounting", "book", "/accounting/payments").SubMenu = func() items {
			var items1 items
			items1.add("payments", "layers", "/accounting/payments").SubMenu = func() items {
				var items2 items
				for _, c := range categories.All() {
					items2.add(c.Name, "search", "/accounting/payments?q=category_id:"+strconv.FormatUint(uint64(c.ID), 10))
				}
				return items2
			}
			items1.add("categories", "tag", "/accounting/categories")
			items1.add("analysis", "", "/accounting/analysis")
			items1.add("balances", "activity", "/accounting/balances")
			items1.add("expenses", "", "/accounting/expenses")
			items1.add("expenses_by_period", "bar-chart-2", "/accounting/expenses/by_period")
			items1.add("expenses_by_category", "pie-chart", "/accounting/expenses/by_category")
			return items1
		}

		firstAssetName := models2.Plural(assets[0])
		items0.add("assets", "box", "/assets/"+firstAssetName).SubMenu = func() items {
			var items1 items
			for _, asset := range assets {
				assetName := models2.Plural(asset)
				items1.add(assetName, asset.Icon(), "/assets/"+assetName)
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

		items0.add("sites", "layout", "/sites").SubMenu = sitesRecursive(sites, 0)

		items0.add("tasks", "check-circle", "/tasks")

		items0.add("settings", "settings", "/settings/users").SubMenu = func() items {
			var items1 items
			items1.add("users", "users", "/settings/users")
			return items1
		}
		return items0
	}}
}

func sitesRecursive(sites repositories.Repository[models2.Site], parentId int) func() items {
	return func() items {
		var items items
		for _, site := range sites.List("parent_id", parentId) {
			add := items.add(site.Name, "file-text", "/sites?id="+strconv.FormatUint(uint64(site.ID), 10))
			add.SubMenu = sitesRecursive(sites, parentId+1)
		}
		return items
	}
}

func (s *navigation) Decorate(req controllers.Request, res controllers.Response) {
	var query string
	if req.URL.RawQuery != "" {
		query += "?" + req.URL.RawQuery
	}
	traverse(res, req.URL.Path+query, s.root(), 0)
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
	var result = &item{}
	maxSimilarity := 0
	for _, n := range navigation {
		if path == n.Path {
			return n
		}
		//if strings.HasPrefix(path, n.Path) {
		//	result = n
		//}
		//if strings.Contains(n.Path, "?") && !strings.Contains(path, "?") {
		//	continue
		//}
		similarity := similarity(n.Path, path)
		if similarity > maxSimilarity {
			result = n
			maxSimilarity = similarity
		}
	}
	return result
}

func similarity(s1, s2 string) int {
	if strings.Contains(s1, "?") {
		return 0
	}

	s1split := strings.Split(s1, "/")
	s2split := strings.Split(s2, "/")

	var count int
	for i := range s1split {
		if len(s2) > i {
			if strings.Split(s1split[i], "?")[0] == strings.Split(s2split[i], "?")[0] {
				count++
			} else {
				break
			}
		}
	}
	return count
}

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
