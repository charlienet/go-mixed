package sets

import (
	"sync"

	"github.com/charlienet/go-mixed/locker"
	"golang.org/x/exp/constraints"
)

type Set[T comparable] interface {
	Add(...T)
	Remove(v T)
	Asc() Set[T]
	Desc() Set[T]
	Contains(T) bool
	ContainsAny(...T) bool
	ContainsAll(...T) bool
	IsEmpty() bool
	ToSlice() []T // 转换为切片
}

var defaultOptions = option{locker: locker.NewEmptyLocker()}

type option struct {
	locker sync.Locker
}

type setFunc func(option)

func WithSync() setFunc {
	return func(o option) {
		o.locker = &sync.RWMutex{}
	}
}

// 并集
func Union[T constraints.Ordered](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewHashSet[T]()
	}
	if len(sets) == 1 {
		return sets[0]
	}

	ret := NewHashSet[T]()
	for i := range sets {
		ret.Add(sets[i].ToSlice()...)
	}

	return ret
}

// 交集
func Intersection[T constraints.Ordered](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewHashSet[T]()
	}
	if len(sets) == 1 {
		return sets[0]
	}

	ret := NewHashSet[T]()
	base := sets[0]
	for _, v := range base.ToSlice() {
		var insert = true
		for _, s := range sets[1:] {
			if !s.Contains(v) {
				insert = false
				break
			}
		}

		if insert {
			ret.Add(v)
		}
	}

	return ret
}

// 差集
func Difference[T constraints.Ordered](main Set[T], sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return main
	}

	ret := NewHashSet[T]()
	for _, v := range sets[0].ToSlice() {
		isDiff := true
		for _, s := range sets {
			if s.Contains(v) {
				isDiff = false
			}
		}

		if isDiff {
			ret.Add(v)
		}
	}

	return ret
}
