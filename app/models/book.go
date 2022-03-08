package models

type Book struct {
	Model
	Name  string
	Sites []Site
}
