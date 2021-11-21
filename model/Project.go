package model

import "Bilance/service/database"

type Project struct {
	Name        string
	Description string
}

func (p Project) Users(database database.Database) []Project {

}
