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

func Standard(log log.Log) Database {
	s := &logAdapter{log}
	path := config.DatabaseLocation()
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{Logger: s})
	if err != nil {
		log.Critical("Failed to connect database")
	}
	return db
}

type logAdapter struct {
	log log.Log
}

func (l *logAdapter) LogMode(level logger.LogLevel) logger.Interface {
	return l
}
func (l *logAdapter) Info(ctx context.Context, msg string, v ...interface{}) {
	l.log.Info(msg, v...)
}
func (l *logAdapter) Warn(ctx context.Context, msg string, v ...interface{}) {
	l.log.Warning(msg, v...)
}
func (l *logAdapter) Error(ctx context.Context, msg string, v ...interface{}) {
	l.log.Error(msg, v...)
}
func (l *logAdapter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	l.log.Debug("SQL %s -> %d", sql, rowsAffected)
	if err != nil {
		l.log.Error(err.Error())
	}
}
