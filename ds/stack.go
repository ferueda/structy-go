package ds

import (
	"errors"
)

type stack[T any] struct {
	stack []T
}

func (s *stack[T]) Push(el T) {
	s.stack = append(s.stack, el)
}

func (s *stack[T]) Pop() (T, error) {
	var r T

	if len(s.stack) == 0 {
		return r, errors.New("empty queue")
	}

	r = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return r, nil
}

func (s *stack[T]) Size() int {
	return len(s.stack)
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{stack: []T{}}
}
