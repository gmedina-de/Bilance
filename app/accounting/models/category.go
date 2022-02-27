package models

import "genuine/core/models"

type Category struct {
	models.Model
	Name  string `class:"form-control" required:"true"`
	Color string `class:"form-control" required:"true" form:"Color,color"`
}
