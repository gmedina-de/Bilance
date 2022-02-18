package main

import (
	"fmt"
	"genuine/core/injector"
)

//
//import (
//	_ "genuine/accounting"
//	_ "genuine/assets"
//	_ "genuine/contacts"
//	_ "genuine/files"
//)

type Log interface {
	Print()
}

type log struct {
}

func newLog() Log {
	return &log{}
}

func (l *log) Print() {
	fmt.Println("SUCCESS")
}

type Database interface {
	Select()
}

type database struct {
	Log Log
}

func newDatabase() Database {
	inject := injector.Inject(&database{})
	inject.Log.Print()
	inject.Select()
	return inject
}

func (d *database) Select() {
	d.Log.Print()
}

func main() {

	injector.Implementations2(newLog)
	db := newDatabase()
	db.Select()

	//core.Init()
}
