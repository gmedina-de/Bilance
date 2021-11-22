package model

import (
	"fmt"
	"time"
)

type Payment struct {
	Id       int
	Name     string
	Amount   EUR
	Date     time.Time
	Picture  string
	Category PaymentCategory
	Tag      *Tag
	Payer    *User
	Payees   *[]User
	Project  *Project
}

type EUR int64

func (m EUR) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f €", x)
}

type PaymentCategory int

const PaymentCategoryExpense = 0
const PaymentCategoryIncome = 1
const PaymentCategoryTransfer = 2
