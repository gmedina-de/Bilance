package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Model[T any] interface {
	FromRequest(request *http.Request, id int64) *T
	FromQuery(row *sql.Rows) *T
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

type EUR int64

func (m EUR) Format() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f", x)
}

func (m EUR) FormatWithCurrency() string {
	return m.Format() + " €"
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

func SumAmounts(payments []Payment) EUR {
	var result EUR
	for _, payment := range payments {
		result += payment.Amount
	}
	return result
}
