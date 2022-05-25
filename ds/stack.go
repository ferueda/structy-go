package ds

import (
	"errors"
)

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(el T) {
	s.stack = append(s.stack, el)
}

func (s *Stack[T]) Pop() (T, error) {
	var r T

	if len(s.stack) == 0 {
		return r, errors.New("empty queue")
	}

	r = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return r, nil
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}
