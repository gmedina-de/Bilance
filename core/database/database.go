package database

type Database interface {
	Migrate(model any)
	Select(result any, query string, params ...any)
	Insert(model any)
	Update(model any)
	Delete(model any)
}
