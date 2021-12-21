package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Model[T any] interface {
	FromRequest(request *http.Request, id int64) *T
	FromQuery(row *sql.Rows) *T
}

func scanAndPanic(row *sql.Rows, dest ...interface{}) {
	err := row.Scan(dest...)
	if err != nil {
		panic(err)
	}
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
