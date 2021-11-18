package service

import (
	"Bilance/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	Create(model interface{})
	RetrieveAll(model interface{})
}

type dbRepository struct {
	db *gorm.DB
}

func DbRepository() Repository {
	// todo: use settings for allowing another databases
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	return &dbRepository{db}
}

func (g *dbRepository) Create(model interface{}) {
	g.db.Table("users")
	g.db.Create(model)
}

func (g *dbRepository) RetrieveAll(model interface{}) {
	g.db.Table("users").Find(model)
}
