package models

type Site struct {
	Model
	Name    string
	Content string
	BookID  uint
}
