package service

import (
	"database/sql"
	"errors"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
	"reflect"
	"strings"
)

type sqliteDatabase struct {
	log Log
	db  *sql.DB
}

func SqliteDatabase(log Log) Database {
	path := "./database.db"
	var db *sql.DB

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		os.Create(path)
		db, _ = sql.Open("sqlite", path)
		ddl, _ := os.ReadFile("./model/model.sql")
		db.Exec(string(ddl))
	} else {
		db, _ = sql.Open("sqlite", path)
	}
	return &sqliteDatabase{log, db}
}

func (s *sqliteDatabase) Select(result interface{}, queryFunction QueryFunc, conditions ...string) {
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

func (s *sqliteDatabase) Insert(model interface{}) sql.Result {
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()
	var columns []string
	var values []string
	for i := 0; i < modelValue.NumField(); i++ {
		field := modelType.Field(i)
		column := field.Name
		// Many-To-Many-Relationships
		if column == "Id" || field.Type.Kind() == reflect.Slice {
			continue
		}
		columns = append(columns, column)
		value := fmt.Sprintf("'%v'", modelValue.Field(i).Interface())
		values = append(values, value)
	}
	return s.execute(`INSERT INTO ` + modelType.Name() + `(` + strings.Join(columns, ",") + `) VALUES (` + strings.Join(values, ",") + `)`)
}

func (s *sqliteDatabase) Update(model interface{}) sql.Result {
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()
	var sets []string
	id := "0"
	for i := 0; i < modelValue.NumField(); i++ {
		field := modelType.Field(i)
		column := field.Name
		value := fmt.Sprintf("'%v'", modelValue.Field(i).Interface())
		// Many-To-Many-Relationships
		if field.Type.Kind() == reflect.Slice {
			continue
		}
		if column == "Id" {
			id = value
		} else {
			sets = append(sets, column+"="+value)
		}
	}
	return s.execute(`UPDATE ` + modelType.Name() + ` SET ` + strings.Join(sets, ",") + ` WHERE Id = ` + id)
}

func (s *sqliteDatabase) Delete(model interface{}) sql.Result {
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()
	id := "0"
	for i := 0; i < modelValue.NumField(); i++ {
		column := modelType.Field(i).Name
		value := fmt.Sprintf("'%v'", modelValue.Field(i).Interface())
		if column == "Id" {
			id = value
		}
	}
	return s.execute(`DELETE FROM ` + modelType.Name() + ` WHERE Id = ` + id)
}

func (s *sqliteDatabase) MultipleDelete(table string, conditions ...string) sql.Result {
	return s.execute(`DELETE FROM ` + table + " " + strings.Join(conditions, " "))
}

func (s *sqliteDatabase) execute(query string) sql.Result {
	s.log.Debug("SQL " + query)
	result, err := s.db.Exec(query)
	if err != nil {
		s.log.Error(err.Error())
	}
	return result
}
