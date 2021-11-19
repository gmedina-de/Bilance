package model

import (
	"Bilance/service/database"
	"log"
	"strings"
)

type User struct {
	Id       int
	Username string
	Password string
}

func RetrieveUsers(db database.Database, conditions ...string) []User {
	var result []User
	row, err := db.Query("SELECT * FROM User " + strings.Join(conditions, " "))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var Id int
		var Username string
		var Password string
		row.Scan(&Id, &Username, &Password)
		result = append(result, User{Id, Username, Password})
	}
	return result
}
