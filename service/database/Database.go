package database

type Database interface {
	Insert(model interface{})
}
