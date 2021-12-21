package model

import (
	"database/sql"
	"net/http"
)

type Category struct {
	Id        int64
	Name      string
	Color     string
	ProjectId int64
}

func (c Category) FromQuery(row *sql.Rows) *Category {
	columns, _ := row.Columns()
	println(columns)
	row.Scan()
	ScanAndPanic(row, &c.Id, &c.Name, &c.Color, &c.ProjectId)
	return &c
}

func (c Category) FromRequest(request *http.Request, id int64) *Category {
	c.Id = id
	c.Name = request.Form.Get("Name")
	c.Color = request.Form.Get("Color")
	c.ProjectId = GetSelectedProjectId(request)
	return &c
}
