package sets

type sorted_set[T comparable] struct {
	sorted []T
	set    Set[T]
}

func NewSortedSet[T comparable]() *sorted_set[T] {
	return &sorted_set[T]{
		set: NewHashSet[T](),
	}
}

func (s *sorted_set[T]) ToSlice() []T {
	return s.sorted
}
