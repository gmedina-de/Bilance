package repository

import (
	"Bilance/service"
	"database/sql"
)

type baseRepository struct {
	database service.Database
}

func (r *baseRepository) Insert(entity interface{}) {
	r.database.Insert(entity)
}

func (r *baseRepository) Update(entity interface{}) {
	r.database.Update(entity)
}

func (r *baseRepository) Delete(entity interface{}) {
	r.database.Delete(entity)
}

func ScanAndPanic(row *sql.Rows, dest ...interface{}) {
	err := row.Scan(dest...)
	if err != nil {
		panic(err)
	}
}
