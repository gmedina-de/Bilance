package model

type Payment struct {
	Id         int64
	Name       string
	Amount     EUR
	Date       Date
	ProjectId  int64
	CategoryId int64
	PayerId    int64
	PayeeId    int64
}
