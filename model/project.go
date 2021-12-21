package model

import (
	"database/sql"
	"net/http"
)

type Project struct {
	Id          int64
	Name        string
	Description string
	UserIds     string
}

func (p Project) FromQuery(row *sql.Rows) *Project {
	ScanAndPanic(row, &p.Id, &p.Name, &p.Description, &p.UserIds)
	return &p
}

func (p Project) FromRequest(request *http.Request, id int64) *Project {
	p.Id = id
	p.Name = request.Form.Get("Name")
	p.Description = request.Form.Get("Description")
	p.UserIds = request.Form.Get("UserIds")
	return &p
}
