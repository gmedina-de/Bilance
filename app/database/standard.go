package database

import (
	"context"
	"flag"
	"genuine/core/log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db = flag.String("db", "./database.db", "database location")

func Standard(log log.Log) Database {
	path := *db
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{Logger: &logAdapter{log: log}})
	if err != nil {
		log.Critical("Failed to connect database")
	}
	return db
}

type logAdapter struct {
	log log.Log
}

func (s *logAdapter) LogMode(level logger.LogLevel) logger.Interface {
	return s
}
func (s *logAdapter) Info(ctx context.Context, msg string, v ...interface{}) {
	s.log.Info(msg, v...)
}
func (s *logAdapter) Warn(ctx context.Context, msg string, v ...interface{}) {
	s.log.Warning(msg, v...)
}
func (s *logAdapter) Error(ctx context.Context, msg string, v ...interface{}) {
	s.log.Error(msg, v...)
}
func (s *logAdapter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	s.log.Debug("SQL %s -> %d", sql, rowsAffected)
	if err != nil {
		s.log.Error(err.Error())
	}
}
