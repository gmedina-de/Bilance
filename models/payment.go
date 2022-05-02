package models

import (
	"fmt"
)

type Payment struct {
	Model
	Name       string   `form:"required"`
	Amount     Currency `form:"required"`
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
	return m.Raw() + " €"
}

func (m Currency) Raw() string {
	x := float64(m)
	x = x / 100
	// todo admit more currencies
	return fmt.Sprintf("%.2f", x)
}

func SumAmounts(payments []Payment) Currency {
	var result Currency
	for _, payment := range payments {
		result += payment.Amount
	}
	return result
}
