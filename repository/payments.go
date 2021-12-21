package repository

import (
	"Bilance/model"
	"Bilance/service"
)

type payments struct {
	generic[model.Payment]
}

func Payments(database service.Database) GRepository[model.Payment] {
	return &payments{generic[model.Payment]{database, model.Payment{}}}
}
