package models

import (
	"fmt"
	"genuine/core/models"
)

//return fmt.Sprintf(`%v<%v%v%v name="%v"%v>%v</%v>`, label, fType, id, class, name, requiredString, value, fType)

type Payment struct {
	Id     int64       `form:"-"`
	Name   string      `class:"form-control" required:"true"`
	Amount EUR         `class:"form-control" required:"true" form:"Amount\" type=\"number\" in=\"0.00\" max=\"100000.00\" step=\"0.01,input"`
	Date   models.Date `class:"form-control" required:"true" form:"Date,date"`
	//CategoryId int64
	Category *Category `class:"form-control" required:"true" form:"Date,date" orm:"rel(fk)"`
	PayerId  int64
	PayeeId  int64
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
