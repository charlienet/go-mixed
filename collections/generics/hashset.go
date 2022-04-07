package generics


type Set[T comparable] map[T]struct{}

func NewHashSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	set.Add(values...)
	return set
}

func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contain(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Clone() Set[T] {
	set := NewHashSet[T]()
	set.Add(s.Values()...)
	return set
}

func (s Set[T]) Iterate(fn func(value T)) {
	for v := range s {
		fn(v)
	}
}

// Union creates a new set contain all element of set s and other
func (s Set[T]) Union(other Set[T]) Set[T] {
	set := s.Clone()
	set.Add(other.Values()...)
	return set
}

// Intersection creates a new set whose element both be contained in set s and other
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	set := NewHashSet[T]()
	s.Iterate(func(value T) {
		if other.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, s.Size())
	s.Iterate(func(value T) {
		values = append(values, value)
	})

	return values
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Size() int {
	return len(s)
}
