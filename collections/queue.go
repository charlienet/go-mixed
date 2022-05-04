package collections

import "sync"

var _ Queue[string] = &ArrayQueue[string]{}

// 数组队列，先进先出
type ArrayQueue[T any] struct {
	array []T        // 底层切片
	size  int        // 队列的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{}
}

// 入队
func (q *ArrayQueue[T]) Put(v T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	q.array = append(q.array, v)

	// 队中元素数量+1
	q.size = q.size + 1
}

// 出队
func (q *ArrayQueue[T]) Poll() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	// 队中元素已空
	if q.size == 0 {
		panic("empty")
	}

	// 队列最前面元素
	v := q.array[0]

	/*    直接原位移动，但缩容后继的空间不会被释放
	      for i := 1; i < queue.size; i++ {
	          // 从第一位开始进行数据移动
	          queue.array[i-1] = queue.array[i]
	      }
	      // 原数组缩容
	      queue.array = queue.array[0 : queue.size-1]
	*/

	// 创建新的数组，移动次数过多
	newArray := make([]T, q.size-1)
	for i := 1; i < q.size; i++ {
		// 从老数组的第一位开始进行数据移动
		newArray[i-1] = q.array[i]
	}
	q.array = newArray

	// 队中元素数量-1
	q.size = q.size - 1
	return v
}

func (q *ArrayQueue[T]) Peek() T {
	return q.array[0]
}

// 栈大小
func (q *ArrayQueue[T]) Size() int {
	return q.size
}

// 栈是否为空
func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.size == 0
}
