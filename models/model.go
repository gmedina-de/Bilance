package models

import "gorm.io/gorm"

type Model = gorm.Model

type Titler interface {
	Title() string
}
