package generics

import "errors"

type ArrayStack[T any] struct {
	data   []T
	length int
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{data: []T{}, length: 0}
}

func (s *ArrayStack[T]) Data() []T {
	return s.data
}

func (s *ArrayStack[T]) Size() int {
	return s.length
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return s.length == 0
}

// Push element into stack
func (s *ArrayStack[T]) Push(value T) {
	s.data = append([]T{value}, s.data...)
	s.length++
}

// Pop delete the top element of stack then return it, if stack is empty, return nil and error
func (s *ArrayStack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	topItem := s.data[0]
	s.data = s.data[1:]
	s.length--

	return &topItem, nil
}

// Peak return the top element of stack then return it
func (s *ArrayStack[T]) Peak() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[0], nil
}

// Clear the stack data
func (s *ArrayStack[T]) Clear() {
	s.data = []T{}
	s.length = 0
}
