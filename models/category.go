package models

type Category struct {
	Model
	Name  string `form:"required"`
	Color string `form:"required"`
}

func (c Category) String() string {
	return "<div style=\"color:" + c.Color + "\">" + c.Name + "</div>"
}

func (c Category) Title() string {
	return c.Name
}
