package assets

import (
	"homecloud/assets/model"
	model2 "homecloud/core/model"
	"homecloud/core/template"
)

func init() {
	menuItem := template.AddNavigation("assets", "box")
	models := model.Models
	for i, m := range models {
		menuItem = menuItem.WithChild(model2.Name(m), model.Icons[i])
	}
	menuItem.Path = "/assets/" + model2.Name(models[0])
}
