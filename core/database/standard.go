package database

import (
	"genuine/core/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type standard struct {
	log log.Log
	db  *gorm.DB
}

func Standard(log log.Log) Database {
	path := "./database.db"
	db, err := gorm.Open(sqlite.Open(path), nil)
	if err != nil {
		log.Critical("Failed to connect database")
	}
	return &standard{log, db}
}

func (s *standard) Migrate(model any) {
	err := s.db.AutoMigrate(model)
	if err != nil {
		s.log.Error(err.Error())
	}
}

func (s *standard) Select(result any, query string, params ...any) {
	s.db.Raw("SELECT "+query, params...).Scan(result)
}

func (s *standard) Insert(model any) {
	s.db.Create(model)
}

func (s *standard) Update(model any) {
	s.db.Save(model)
}

func (s *standard) Delete(model any) {
	s.db.Delete(model)
}
