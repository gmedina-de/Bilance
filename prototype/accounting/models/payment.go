package models

import (
	"fmt"
	"genuine/framework/models"
)

type Payment struct {
	Id         int64
	Name       string
	Amount     EUR
	Date       models.Date
	CategoryId int64
	PayerId    int64
	PayeeId    int64
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

func SumAmounts(payments []Payment) EUR {
	var result EUR
	for _, payment := range payments {
		result += payment.Amount
	}
	return result
}
