package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type standard struct {
	db *gorm.DB
}

func Standard() Database {
	path := "./database.db"
	db, err := gorm.Open(sqlite.Open(path), nil)
	if err != nil {
		panic("failed to connect database")
	}
	return &standard{db}
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
