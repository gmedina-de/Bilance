package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Model any

func NamePlural(model any) string {
	name := Name(model)
	if name[len(name)-1] == 'y' {
		return name[0:len(name)-1] + "ies"
	}
	return name + "s"
}

func Name(model any) string {
	return strings.ToLower(reflect.TypeOf(model).Name())
}

func ScanAndPanic(row *sql.Rows, dest ...interface{}) {
	err := row.Scan(dest...)
	if err != nil {
		panic(err)
	}
}

func StringToIds(input string) []int64 {
	split := strings.Split(input, ",")
	var result []int64
	for _, id := range split {
		parseInt, _ := strconv.ParseInt(id, 10, 64)
		result = append(result, parseInt)
	}
	return result
}

func Serialize(T any) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(T)
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func Deserialize[T any](str string, dest *T) {
	by, _ := base64.StdEncoding.DecodeString(str)
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	_ = d.Decode(&dest)
}

type Date string

func (d Date) Format() string {
	parse, _ := time.Parse(DateLayoutISO, string(d))
	return parse.Format(DateLayoutDE)
}

const (
	DateLayoutISO = "2006-01-02"
	DateLayoutDE  = "02.01.2006"
)

type TimeUnit int

const TimeUnitWeekday TimeUnit = iota
const TimeUnitMonthday TimeUnit = iota + 1
const TimeUnitMonth TimeUnit = iota + 2

func NormalWeekday(weekday time.Weekday) int {
	if weekday == time.Sunday {
		return 6
	}
	return int(weekday) - 1
}
