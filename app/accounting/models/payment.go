package models

import (
	"fmt"
	models2 "genuine/app/common/models"
	"genuine/core/models"
)

type Payment struct {
	Id       int64    `form:"-"`
	Name     string   `form:"required"`
	Amount   Currency `form:"required"`
	Date     models.Date
	Category Category
	Payer    models2.User
	Payee    models2.User
}

type Currency int64

func (m Currency) Format() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f", x)
}

func (m Currency) FormatWithCurrency() string {
	return m.Format() + " €"
}

func SumAmounts(payments []Payment) Currency {
	var result Currency
	for _, payment := range payments {
		result += payment.Amount
	}
	return result
}
