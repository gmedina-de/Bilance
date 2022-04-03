package models

import (
	"fmt"
)

type Payment struct {
	Model
	Name       string   `form:"required"`
	Amount     Currency `form:"required min=\"0.00\" max=\"100000.00\" step=\"0.01\""`
	Date       Date
	CategoryID uint
	Category   Category
	PayerID    uint
	Payer      User
	PayeeID    uint
	Payee      User
}

type Currency int64

func (m Currency) String() string {
	x := float64(m)
	x = x / 100
	// todo admit more currencies
	return fmt.Sprintf("%.2f", x) + " €"
}

func SumAmounts(payments []Payment) Currency {
	var result Currency
	for _, payment := range payments {
		result += payment.Amount
	}
	return result
}
