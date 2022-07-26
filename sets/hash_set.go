package sets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charlienet/go-mixed/locker"
	"golang.org/x/exp/constraints"
)

var _ Set[string] = &hash_set[string]{}

type hash_set[T constraints.Ordered] struct {
	m    map[T]struct{}
	lock locker.RWLocker
}

func NewHashSet[T constraints.Ordered](values ...T) *hash_set[T] {
	set := hash_set[T]{
		m:    make(map[T]struct{}, len(values)),
		lock: locker.EmptyLocker,
	}

	set.Add(values...)
	return &set
}

func (s *hash_set[T]) WithSync() *hash_set[T] {
	s.lock = locker.NewRWLocker()
	return s
}

func (s hash_set[T]) Add(values ...T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, v := range values {
		s.m[v] = struct{}{}
	}
}

func (s hash_set[T]) Remove(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.m, v)
}

func (s hash_set[T]) Contains(value T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.m[value]
	return ok
}

func (s hash_set[T]) ContainsAny(values ...T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, v := range values {
		if _, ok := s.m[v]; ok {
			return true
		}
	}

	return false
}

func (s hash_set[T]) ContainsAll(values ...T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, v := range values {
		if _, ok := s.m[v]; !ok {
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
	for k := range s.m {
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
	for v := range s.m {
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
	return len(s.m) == 0
}

func (s hash_set[T]) Size() int {
	return len(s.m)
}

func (s hash_set[T]) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Size())

	for ele := range s.m {
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
	l := make([]string, 0, len(s.m))
	for k := range s.m {
		l = append(l, fmt.Sprint(k))
	}

	return fmt.Sprintf("{%s}", strings.Join(l, ", "))
}
