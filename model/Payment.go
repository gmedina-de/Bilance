package model

import (
	"fmt"
	"time"
)

type Payment struct {
	Name     string
	Amount   EUR
	Date     time.Time
	Picture  string
	Category Category
	Tag      *Tag
	Payer    *User
	Payee    *User
	Project  *Project
}

type EUR int64

func (m EUR) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f €", x)
}

type Category int

const CategoryExpense = 0
const CategoryIncome = 1
const CategoryTransfer = 2
