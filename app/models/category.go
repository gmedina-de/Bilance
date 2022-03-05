package models

type Category struct {
	Model
	Name  string `form:"required"`
	Color string `form:"required"`
}

func (c Category) String() string {
	return c.Name
}
