package generics

type Map[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Delete(key K)
	ForEach(f func(K, V))
	Clone() Map[K, V]
	Clear()
	Count() int
}
