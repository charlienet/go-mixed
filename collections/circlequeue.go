package collections

import "fmt"

type CircleQueue struct {
	data  []interface{}
	cap   int
	front int
	rear  int
}

func NewCircleQueue(cap int) *CircleQueue {
	cap++

	return &CircleQueue{
		data: make([]interface{}, cap),
		cap:  cap,
	}
}

func (q *CircleQueue) Push(data interface{}) bool {
	if (q.rear+1)%q.cap == q.front {
		return false
	}

	q.data[q.rear] = data
	q.rear = (q.rear + 1) % q.cap
	return true
}
func (q *CircleQueue) Pop() interface{} {
	if q.rear == q.front {
		return nil
	}

	data := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.cap
	return data
}

func (q *CircleQueue) Size() int {
	return (q.rear - q.front + q.cap) % q.cap
}

func (q *CircleQueue) IsFull() bool {
	return (q.rear+1)%q.cap == q.front
}

func (q *CircleQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *CircleQueue) Show() string {
	return fmt.Sprintf("%v", q.data)
}
