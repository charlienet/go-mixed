package sets

import (
	"slices"

	"github.com/charlienet/go-mixed/expr"
	"golang.org/x/exp/constraints"
)

type sorted_set[T constraints.Ordered] struct {
	sorted []T
	set    Set[T]
}

func NewSortedSet[T constraints.Ordered](t ...T) *sorted_set[T] {
	return &sorted_set[T]{
		sorted: t,
		set:    NewHashSet(t...),
	}
}

func (s *sorted_set[T]) Add(values ...T) Set[T] {
	for _, v := range values {
		if !s.set.Contains(v) {
			s.sorted = append(s.sorted, v)
			s.set.Add(v)
		}
	}

	return s
}

func (s *sorted_set[T]) Remove(v T) Set[T] {
	if s.set.Contains(v) {
		for index := range s.sorted {
			if s.sorted[index] == v {
				s.sorted = append(s.sorted[:index], s.sorted[index+1:]...)
				break
			}
		}

		s.set.Remove(v)
	}

	return s
}

func (s *sorted_set[T]) Asc() Set[T] {
	keys := s.sorted
	slices.Sort(keys)

	return &sorted_set[T]{
		sorted: keys,
		set:    NewHashSet(keys...),
	}
}

func (s *sorted_set[T]) Desc() Set[T] {
	keys := s.sorted
	slices.SortFunc(keys, func(a, b T) int {
		if a == b {
			return 0
		}

		return expr.Ternary(a > b, -1, 1)
	})

	return &sorted_set[T]{
		sorted: keys,
		set:    NewHashSet(keys...),
	}
}

func (s *sorted_set[T]) Contains(v T) bool {
	return s.set.Contains(v)
}

func (s *sorted_set[T]) ContainsAny(values ...T) bool {
	return s.set.ContainsAny(values...)
}

func (s *sorted_set[T]) ContainsAll(values ...T) bool {
	return s.set.ContainsAll(values...)
}

func (s *sorted_set[T]) IsEmpty() bool {
	return s.set.IsEmpty()
}

func (s *sorted_set[T]) ToSlice() []T {
	return s.sorted
}
