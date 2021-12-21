package repository

import (
	"Bilance/database"
	"Bilance/model"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

type payments struct {
	generic[model.Payment]
}

func Payments(database database.Database) Repository[model.Payment] {
	return &payments{
		generic[model.Payment]{
			database: database,
			model:    model.Payment{},
		},
	}
}

func (p *payments) FromQuery(row *sql.Rows) *model.Payment {
	payment := model.Payment{}
	model.ScanAndPanic(row, &payment.Id, &payment.Name, &payment.Amount, &payment.Date, &payment.ProjectId, &payment.CategoryId, &payment.PayerId, &payment.PayeeId)
	return &payment
}
func (p *payments) FromRequest(request *http.Request, id int64) *model.Payment {
	payment := model.Payment{}
	payment.Id = id
	payment.Name = request.Form.Get("Name")
	payment.Date = model.Date(request.Form.Get("Date"))
	payment.ProjectId = model.GetSelectedProjectId(request)
	amountInput := request.Form.Get("Amount")
	amountString := strings.Replace(amountInput, ".", "", 1)
	amount, _ := strconv.ParseInt(amountString, 10, 64)
	if !strings.Contains(amountInput, ".") {
		amount *= 100
	}
	payment.Amount = model.EUR(amount)
	payment.CategoryId, _ = strconv.ParseInt(request.Form.Get("CategoryId"), 10, 64)
	payment.PayerId, _ = strconv.ParseInt(request.Form.Get("PayerId"), 10, 64)
	payment.PayeeId, _ = strconv.ParseInt(request.Form.Get("PayeeId"), 10, 64)
	return &payment
}
