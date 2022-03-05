package models

import (
	"genuine/core/models"
)

type Category struct {
	models.Model
	Name  string `form:"required"`
	Color string `form:"required"`
}

func (c Category) String() string {
	return c.Name
}
