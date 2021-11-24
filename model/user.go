package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Role     UserRole
}

type UserRole int64

const UserRoleNormal UserRole = 0
const UserRoleAdmin UserRole = 1

func (user *User) Serialize() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(user)
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func DeserializeUser(str string) *User {
	user := User{}
	by, _ := base64.StdEncoding.DecodeString(str)
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	_ = d.Decode(&user)
	return &user
}
