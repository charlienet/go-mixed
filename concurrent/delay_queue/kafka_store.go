package delayqueue

import (
	"context"

	"github.com/charlienet/go-mixed/errors"
)

type kafkaStore[T Delayed] struct {
}

func (s *delayQueue[T]) UseKafka() *delayQueue[T] {
	s.UseStore(newKafka[T]())

	panic(errors.NotImplemented)
	// return s.UseStore(newKafka[T]())
}

func newKafka[T Delayed]() *kafkaStore[T] {
	return &kafkaStore[T]{}
}

func (*kafkaStore[T]) Push(context.Context, T) error {
	return nil
}

func (*kafkaStore[T]) Pop() (T, error) {
	return *new(T), nil
}

func (*kafkaStore[T]) Peek() (T, bool) {
	return *new(T), false
}

func (*kafkaStore[T]) IsEmpty() bool {
	return false
}
