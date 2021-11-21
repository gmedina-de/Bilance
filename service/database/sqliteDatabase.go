package database

import (
	"Bilance/service/log"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"reflect"
	"strings"
)

type sqliteDatabase struct {
	log log.Log
	db  *sql.DB
}

func SqliteDatabase(log log.Log) Database {
	db, _ := sql.Open("sqlite", "./database.db")
	return &sqliteDatabase{log, db}
}

func (s *sqliteDatabase) Query(result interface{}, queryFunction QueryFunc, conditions ...string) {
	modelType := reflect.TypeOf(result)
	resultValue := reflect.ValueOf(result)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
		resultValue = resultValue.Elem()
	}
	if modelType.Kind() == reflect.Slice {
		modelType = modelType.Elem()
	} else {
		panic("Input param is not a slice")
	}

	query := "SELECT * FROM " + modelType.Name() + " " + strings.Join(conditions, " ")
	s.log.Debug("SQL " + query)
	row, err := s.db.Query(query)
	if err != nil {
		s.log.Error(err.Error())
	}
	defer row.Close()
	for row.Next() {
		resultValue.Set(reflect.Append(resultValue, reflect.ValueOf(queryFunction(row)).Elem()))
	}
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
	s.Exec(`INSERT INTO ` + modelType.Name() + `(` + strings.Join(columns, ",") + `) VALUES (` + strings.Join(values, ",") + `)`)
}

func (s *sqliteDatabase) Update(model interface{}) {
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()
	var sets []string
	id := "0"
	for i := 0; i < modelValue.NumField(); i++ {
		column := modelType.Field(i).Name
		value := fmt.Sprintf("'%v'", modelValue.Field(i).Interface())
		if column == "Id" {
			id = value
		} else {
			sets = append(sets, column+"="+value)
		}
	}
	s.Exec(`UPDATE ` + modelType.Name() + ` SET ` + strings.Join(sets, ",") + ` WHERE Id = ` + id)
}

func (s *sqliteDatabase) Delete(table string, id string) {
	s.Exec(`DELETE FROM ` + table + ` WHERE Id = '` + id + `'`)
}

func (s *sqliteDatabase) Exec(query string) {
	s.log.Debug("SQL " + query)
	_, err := s.db.Exec(query)
	if err != nil {
		s.log.Error(err.Error())
	}
}
