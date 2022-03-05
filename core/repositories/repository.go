package repositories

type Repository[T any] interface {
	All() []T
	Count(where string, args ...any) int64
	Find(id uint) *T
	Limit(limit int, offset int, where string, args ...any) []T
	List(where string, args ...any) []T

	Insert(entity *T)
	Update(entity *T)
	Delete(entity *T)

	Model() T
}
