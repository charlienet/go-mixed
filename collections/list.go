package collections

type List[T any] interface {
	Add(T)
	Delete(T)
	Count() int
	ToSlice() []T
}


type Queue interface{
	
}