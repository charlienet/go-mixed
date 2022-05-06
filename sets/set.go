package sets

type Set[T comparable] interface {
	Add(values ...T)
	Remove(v T)
	Contain(value T) bool
	IsEmpty() bool
}
