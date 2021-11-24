package model

import (
	"fmt"
)

type Payment struct {
	Id        int64
	Name      string
	Amount    EUR
	Date      Date
	ProjectId int64
	Type      *Type
	Payer     *User
	Payee     *User
}

type EUR int64

func (m EUR) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f", x)
}

type Date string
