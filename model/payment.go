package model

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

type Payment struct {
	Id         int64
	Name       string
	Amount     EUR
	Date       Date
	ProjectId  int64
	CategoryId int64
	PayerId    int64
	PayeeId    int64
}

func (p Payment) FromQuery(row *sql.Rows) *Payment {
	ScanAndPanic(row, &p.Id, &p.Name, &p.Amount, &p.Date, &p.ProjectId, &p.CategoryId, &p.PayerId, &p.PayeeId)
	return &p
}

func (p Payment) FromRequest(request *http.Request, id int64) *Payment {
	p.Id = id
	p.Name = request.Form.Get("Name")
	p.Date = Date(request.Form.Get("Date"))
	p.ProjectId = GetSelectedProjectId(request)
	amountInput := request.Form.Get("Amount")
	amountString := strings.Replace(amountInput, ".", "", 1)
	amount, _ := strconv.ParseInt(amountString, 10, 64)
	if !strings.Contains(amountInput, ".") {
		amount *= 100
	}
	p.Amount = EUR(amount)
	p.CategoryId, _ = strconv.ParseInt(request.Form.Get("CategoryId"), 10, 64)
	p.PayerId, _ = strconv.ParseInt(request.Form.Get("PayerId"), 10, 64)
	p.PayeeId, _ = strconv.ParseInt(request.Form.Get("PayeeId"), 10, 64)
	return &p
}
