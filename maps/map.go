package maps

import "golang.org/x/exp/constraints"

type Map[K constraints.Ordered, V any] interface {
	Set(key K, value V)           // 设置值
	Get(key K) (value V, ok bool) // 获取值
	Exist(key K) bool             // 键是否存在
	Delete(key K)                 // 删除值
	Keys() []K                    // 获取所有键
	Values() []V                  // 获取所有值
	ToMap() map[K]V               // 转换为map
	Clone() Map[K, V]             // 复制
	Clear()                       // 清空
	Count() int                   // 数量
	Iter() <-chan *Entry[K, V]
	ForEach(f func(K, V))
}

type Entry[K constraints.Ordered, V any] struct {
	Key   K
	Value V
}

func Merge[K comparable, V any](mm ...map[K]V) map[K]V {
	ret := make(map[K]V)
	for _, m := range mm {
		for k, v := range m {
			ret[k] = v
		}
	}

	return ret
}
