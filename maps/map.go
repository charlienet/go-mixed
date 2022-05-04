package maps

import "golang.org/x/exp/constraints"

type Map[K constraints.Ordered, V any] interface {
	Set(key K, value V)
	Get(key K) (value V, ok bool)
	Exist(key K) bool
	Delete(key K)
	Clone() Map[K, V]
	Clear()
	Count() int
	ForEach(f func(K, V))
}
