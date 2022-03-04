package repositories

type Repository[T any] interface {
	All() []T
	Count() int64
	Find(id uint) *T
	Limit(limit int, offset int) []T
	List(where string, args ...any) []T

	Insert(entity *T)
	Update(entity *T)
	Delete(entity *T)

	Model() T
}
