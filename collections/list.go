package collections

// 列表
type List[T any] interface {
	Add(T)
	Delete(T)
	Count() int
	ToSlice() []T
}

// 队列
type Queue[T any] interface {
	Put(T)
	Poll() T
	Peek() T
	Size() int
	IsEmpty() bool
}
