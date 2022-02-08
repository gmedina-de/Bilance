package repositories

import (
	"genuine/accounting/models"
	"genuine/core/database"
	"genuine/core/repositories"
)

type payments struct {
	repositories.Repository[models.Payment]
}

func Payments(database database.Database) repositories.Repository[models.Payment] {
	return &payments{repositories.Generic(database, models.Payment{})}
}

//func (p *payments) FromRequest(request *http.Request, id int64) *models.Payment {
//	payment := models.Payment{}
//	payment.Id = id
//	payment.Name = request.Form.Get("Name")
//	payment.Date = model2.Date(request.Form.Get("Date"))
//	amountInput := request.Form.Get("Amount")
//	amountString := strings.Replace(amountInput, ".", "", 1)
//	amount, _ := strconv.ParseInt(amountString, 10, 64)
//	if !strings.Contains(amountInput, ".") {
//		amount *= 100
//	}
//	payment.Amount = models.EUR(amount)
//	payment.CategoryId, _ = strconv.ParseInt(request.Form.Get("CategoryId"), 10, 64)
//	payment.PayerId, _ = strconv.ParseInt(request.Form.Get("PayerId"), 10, 64)
//	payment.PayeeId, _ = strconv.ParseInt(request.Form.Get("PayeeId"), 10, 64)
//	return &payment
//}
