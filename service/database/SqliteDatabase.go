package database

import (
	"Bilance/model"
	"Bilance/service/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDatabase struct {
	log log.Log
	db  *gorm.DB
}

func SqliteDatabase(log log.Log) Database {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	return &sqliteDatabase{log, db}
}

func (g *sqliteDatabase) Create(model interface{}) {
	modelName := modelName(model)
	g.log.Info("Creating a " + modelName)
	g.db.Table(modelName).Create(model)
}

func (g *sqliteDatabase) RetrieveAll(model interface{}) {
	modelName := modelName(model)
	g.log.Info("Retrieving all " + modelName)
	g.db.Table(modelName).Find(model)
}
