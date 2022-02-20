package assets

import (
	"genuine/apps/assets/models"
	model2 "genuine/core/models"
	"genuine/core/template"
)

func init() {
	menuItem := template.AddNavigation("assets", "box")
	Models := models.Models
	for i, m := range Models {
		menuItem = menuItem.WithChild(model2.Plural(m), models.Icons[i])
	}
	menuItem.Path = "/assets/" + model2.Plural(Models[0])
}
