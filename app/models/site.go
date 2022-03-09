package models

type Site struct {
	Model
	Name     string
	Content  string
	ParentID uint
	Parent   *Site
}
