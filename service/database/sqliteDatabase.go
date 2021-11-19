package database

import (
	"Bilance/service/log"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"strings"
)

type sqliteDatabase struct {
	log log.Log
	db  *sql.DB
}

func SqliteDatabase(log log.Log) Database {
	db, _ := sql.Open("sqlite3", "./database.db")
	return &sqliteDatabase{log, db}
}

func (s *sqliteDatabase) Insert(model interface{}) {
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()
	var columns []string
	var values []string
	for i := 0; i < modelValue.NumField(); i++ {
		column := modelType.Field(i).Name
		if column != "Id" {
			columns = append(columns, column)
			value := fmt.Sprintf("'%v'", modelValue.Field(i).Interface())
			values = append(values, value)
		}
	}

	query := `INSERT INTO ` + modelType.Name() + `(` + strings.Join(columns, ",") + `) VALUES (` + strings.Join(values, ",") + `)`
	s.log.Debug("Executing query " + query)
	_, err := s.db.Exec(query)
	if err != nil {
		s.log.Error(err.Error())
	}
}

func (s *sqliteDatabase) Query(query string) (*sql.Rows, error) {
	return s.db.Query(query)
}
