package sets

type hash_set[T comparable] map[T]struct{}

func NewHashSet[T comparable](values ...T) hash_set[T] {
	set := make(hash_set[T], len(values))
	set.Add(values...)
	return set
}

func (s hash_set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

func (s hash_set[T]) Remove(v T) {
	delete(s, v)
}

func (s hash_set[T]) Contain(value T) bool {
	_, ok := s[value]
	return ok
}

func (s hash_set[T]) Clone() hash_set[T] {
	set := NewHashSet[T]()
	set.Add(s.Values()...)
	return set
}

func (s hash_set[T]) Iterate(fn func(value T)) {
	for v := range s {
		fn(v)
	}
}

// Union creates a new set contain all element of set s and other
func (s hash_set[T]) Union(other hash_set[T]) hash_set[T] {
	set := s.Clone()
	set.Add(other.Values()...)
	return set
}

// Intersection creates a new set whose element both be contained in set s and other
func (s hash_set[T]) Intersection(other hash_set[T]) hash_set[T] {
	set := NewHashSet[T]()
	s.Iterate(func(value T) {
		if other.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

func (s hash_set[T]) Values() []T {
	values := make([]T, 0, s.Size())
	s.Iterate(func(value T) {
		values = append(values, value)
	})

	return values
}

func (s hash_set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s hash_set[T]) Size() int {
	return len(s)
}
