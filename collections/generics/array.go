package generics

type Array[T comparable] struct {
	data   []T
	length int
}

func NewArray[T comparable](values ...T) *Array[T] {
	return &Array[T]{data: values, length: len(values)}
}

func (a *Array[T]) Distinct(keepSorted bool) *Array[T] {
	set := NewHashSet(a.data...)
	if set.Size() == a.length {
		return a
	}

	if !keepSorted {
		return NewArray(set.Values()...)
	}

	ret := make([]T, 0, len(set))
	for _, v := range a.data {
		if set.Contain(v) {
			ret = append(ret, v)
			set.Remove(v)
		}
	}

	return NewArray(ret...)
}

func (a *Array[T]) ToList() []T {
	return a.data
}
