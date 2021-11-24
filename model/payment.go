package model

import (
	"fmt"
	"time"
)

type Payment struct {
	Id        int64
	Name      string
	Amount    EUR
	Date      time.Time
	ProjectId int64
	Tag       *Tag
	Payer     *User
	Payee     *User
}

type EUR int64

func (m EUR) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f €", x)
}
