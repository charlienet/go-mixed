package sets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

var _ Set[string] = &hash_set[string]{}

type hash_set[T constraints.Ordered] map[T]struct{}

func NewHashSet[T constraints.Ordered](values ...T) *hash_set[T] {
	set := make(hash_set[T], len(values))
	set.Add(values...)
	return &set
}

func (s hash_set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

func (s hash_set[T]) Remove(v T) {
	delete(s, v)
}

func (s hash_set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s hash_set[T]) ContainsAny(values ...T) bool {
	for _, v := range values {
		if _, ok := s[v]; ok {
			return true
		}
	}

	return false
}

func (s hash_set[T]) ContainsAll(values ...T) bool {
	for _, v := range values {
		if _, ok := s[v]; !ok {
			return false
		}
	}

	return true
}

func (s hash_set[T]) Asc() Set[T] {
	return s.copyToSorted().Asc()
}

func (s hash_set[T]) Desc() Set[T] {
	return s.copyToSorted().Desc()
}

func (s hash_set[T]) copyToSorted() Set[T] {
	orderd := NewSortedSet[T]()
	for k := range s {
		orderd.Add(k)
	}

	return orderd
}

func (s *hash_set[T]) Clone() *hash_set[T] {
	set := NewHashSet[T]()
	set.Add(s.ToSlice()...)
	return set
}

func (s hash_set[T]) Iterate(fn func(value T)) {
	for v := range s {
		fn(v)
	}
}

func (s hash_set[T]) ToSlice() []T {
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

func (s hash_set[T]) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Size())

	for ele := range s {
		b, err := json.Marshal(ele)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ", "))), nil
}

func (s hash_set[T]) UnmarshalJSON(b []byte) error {
	var i []any

	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}

	for _, v := range i {
		if t, ok := v.(T); ok {
			s.Add(t)
		}
	}

	return nil
}

func (s hash_set[T]) String() string {
	l := make([]string, 0, len(s))
	for k := range s {
		l = append(l, fmt.Sprint(k))
	}

	return fmt.Sprintf("{%s}", strings.Join(l, ", "))
}
