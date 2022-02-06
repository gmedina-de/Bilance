package assets

import (
	"homecloud/assets/models"
	model2 "homecloud/core/models"
	"homecloud/core/template"
)

func init() {
	menuItem := template.AddNavigation("assets", "box")
	Models := models.Models
	for i, m := range Models {
		menuItem = menuItem.WithChild(model2.Plural(m), models.Icons[i])
	}
	menuItem.Path = "/assets/" + model2.Plural(Models[0])
}
