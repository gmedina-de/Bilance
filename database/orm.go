package database

//
//import (
//	"Bilance/log"
//	"database/sql"
//	"errors"
//	"fmt"
//	_ "github.com/mattn/go-sqlite3"
//	"os"
//	"reflect"
//	"strings"
//)
//
//type orm struct {
//	log log.Log
//	db  *sql.DB
//}
//
//func ORM(log log.Log) Database {
//	path := "./database.db"
//	var db *sql.DB
//
//	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
//		os.Create(path)
//		db, _ = sql.Open("sqlite", path)
//		ddl, _ := os.ReadFile("./model/model.sql")
//		db.Exec(string(ddl))
//		db.Exec("insert into User values (null,'admin','admin',1);")
//	} else {
//		db, _ = sql.Open("sqlite", path)
//	}
//	return &orm{log, db}
//}
//
//func (s *orm) Select(table string, result interface{}, columns string, queryFunction QueryFunc, conditions ...string) {
//	modelType := reflect.TypeOf(result)
//	resultValue := reflect.ValueOf(result)
//	if modelType.Kind() == reflect.Ptr {
//		modelType = modelType.Elem()
//		resultValue = resultValue.Elem()
//	}
//	if modelType.Kind() == reflect.Slice {
//		modelType = modelType.Elem()
//	} else {
//		panic("Input param is not a slice")
//	}
//
//	query := "SELECT " + columns + " FROM " + table + " " + strings.Join(conditions, " ")
//	s.log.Debug("SQL " + query)
//	row, err := s.db.Query(query)
//	if err != nil {
//		s.log.Error(err.Error())
//	}
//	defer row.Close()
//	for row.Next() {
//		resultValue.Set(reflect.Append(resultValue, reflect.ValueOf(queryFunction(row)).Elem()))
//	}
//}
//
//func (s *orm) Insert(table string, model interface{}) sql.Result {
//	modelType := reflect.TypeOf(model).Elem()
//	modelValue := reflect.ValueOf(model).Elem()
//	var columns []string
//	var values []string
//	for i := 0; i < modelValue.NumField(); i++ {
//		field := modelType.Field(i)
//		column := field.Name
//		value := modelValue.Field(i)
//
//		// Many-To-Many-Relationships
//		if column == "Id" || field.Type.Kind() == reflect.Slice {
//			continue
//		}
//
//		// One-To-Many-Relationships
//		if field.Type.Kind() == reflect.Ptr {
//			column += "Id"
//			value = value.Elem().FieldByName("Id")
//		}
//
//		columns = append(columns, column)
//		values = append(values, fmt.Sprintf("'%v'", value.Interface()))
//	}
//	return s.execute(`INSERT INTO ` + table + `(` + strings.Join(columns, ",") + `) VALUES (` + strings.Join(values, ",") + `)`)
//}
//
//func (s *orm) Update(table string, model interface{}) sql.Result {
//	modelType := reflect.TypeOf(model).Elem()
//	modelValue := reflect.ValueOf(model).Elem()
//	var sets []string
//	id := "0"
//	for i := 0; i < modelValue.NumField(); i++ {
//		field := modelType.Field(i)
//		column := field.Name
//		value := modelValue.Field(i)
//
//		// Many-To-Many-Relationships
//		if field.Type.Kind() == reflect.Slice {
//			continue
//		}
//
//		// One-To-Many-Relationships
//		if field.Type.Kind() == reflect.Ptr {
//			column += "Id"
//			value = value.Elem().FieldByName("Id")
//		}
//
//		valueString := fmt.Sprintf("'%v'", value.Interface())
//		if column == "Id" {
//			id = valueString
//		} else {
//			sets = append(sets, column+"="+valueString)
//		}
//	}
//	return s.execute(`UPDATE ` + table + ` SET ` + strings.Join(sets, ",") + ` WHERE Id = ` + id)
//}
//
//func (s *orm) Delete(table string, model interface{}) sql.Result {
//	modelType := reflect.TypeOf(model).Elem()
//	modelValue := reflect.ValueOf(model).Elem()
//	id := "0"
//	for i := 0; i < modelValue.NumField(); i++ {
//		column := modelType.Field(i).Name
//		value := fmt.Sprintf("'%v'", modelValue.Field(i).Interface())
//		if column == "Id" {
//			id = value
//		}
//	}
//	return s.execute(`DELETE FROM ` + table + ` WHERE Id = ` + id)
//}
//
//func (s *orm) MultipleDelete(table string, conditions ...string) sql.Result {
//	return s.execute(`DELETE FROM ` + table + " " + strings.Join(conditions, " "))
//}
//
//func (s *orm) execute(query string) sql.Result {
//	s.log.Debug("SQL " + query)
//	result, err := s.db.Exec(query)
//	if err != nil {
//		s.log.Error(err.Error())
//	}
//	return result
//}
