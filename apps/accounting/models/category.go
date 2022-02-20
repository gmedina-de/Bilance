package models

type Category struct {
	Id    int64  `form:"-"`
	Name  string `class:"form-control" required:"true"`
	Color string `class:"form-control" required:"true" form:"Color,color"`
}
