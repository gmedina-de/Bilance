package models

import (
	"fmt"
	"genuine/core/models"
)

type Payment struct {
	models.Model
	Name       string   `form:"required"`
	Amount     Currency `form:"required"`
	Date       models.Date
	CategoryID uint
	Category   Category
	PayerID    uint
	Payer      User
	PayeeID    uint
	Payee      User
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
