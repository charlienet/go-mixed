package maps

import (
	"strings"

	"golang.org/x/exp/constraints"
)

type hashable interface {
	constraints.Integer | constraints.Float | ~string
}

type Map[K hashable, V any] interface {
	Set(key K, value V)           // 设置值
	Get(key K) (value V, ok bool) // 获取值
	Exist(key K) bool             // 键是否存在
	Delete(key K)                 // 删除值
	Keys() []K                    // 获取所有键
	Values() []V                  // 获取所有值
	ToMap() map[K]V               // 转换为map
	Count() int                   // 数量
	// Iter() <-chan *Entry[K, V]    // 迭代器
	ForEach(f func(K, V) bool) // ForEach
}

type Entry[K hashable, V any] struct {
	Key   K
	Value V
}

func Merge[K hashable, V any](mm ...map[K]V) map[K]V {
	ret := make(map[K]V)
	for _, m := range mm {
		for k, v := range m {
			ret[k] = v
		}
	}

	return ret
}

// 按照键值生成字符串
func Join[K hashable, V any](m Map[K, V], sep string, f func(k K, v V) string) string {
	slice := make([]string, 0, m.Count())

	m.ForEach(func(k K, v V) bool {
		slice = append(slice, f(k, v))
		return false
	})

	return strings.Join(slice, sep)
}
