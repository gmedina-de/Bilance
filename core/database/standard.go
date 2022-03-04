package database

import (
	"context"
	"genuine/core/config"
	"genuine/core/log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type standard struct {
	log log.Log
	db  *gorm.DB
}

func Standard(log log.Log) Database {
	s := &standard{log, nil}
	path := config.DatabaseLocation()
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{Logger: s})
	if err != nil {
		log.Critical("Failed to connect database")
	}
	s.db = db
	return s
}

func (s *standard) Migrate(model any) {
	err := s.db.AutoMigrate(&model)
	if err != nil {
		s.log.Error(err.Error())
	}
}

func (s *standard) Select(result any, query string, params ...any) {
	s.db.Raw("SELECT "+query, params...).Where("deleted_at = NULddddL").Scan(result)
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

// log methods for gorm
func (s *standard) LogMode(level logger.LogLevel) logger.Interface {
	return s
}
func (s *standard) Info(ctx context.Context, msg string, v ...interface{}) {
	s.log.Info(msg, v...)
}
func (s *standard) Warn(ctx context.Context, msg string, v ...interface{}) {
	s.log.Warning(msg, v...)
}
func (s *standard) Error(ctx context.Context, msg string, v ...interface{}) {
	s.log.Error(msg, v...)
}
func (s *standard) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	s.log.Debug("SQL %s -> %d", sql, rowsAffected)
	if err != nil {
		s.log.Error(err.Error())
	}
}
