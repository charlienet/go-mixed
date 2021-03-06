package collections

// ๅ่กจ
type List[T any] interface {
	Add(T)
	Delete(T)
	Count() int
	ToSlice() []T
}

// ้ๅ
type Queue[T any] interface {
	Put(T)
	Poll() T
	Peek() T
	Size() int
	IsEmpty() bool
}
