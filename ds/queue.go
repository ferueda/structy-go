package ds

import (
	"errors"
)

type queue[T any] struct {
	queue []T
}

func (q *queue[T]) Enqueue(el T) {
	q.queue = append(q.queue, el)
}

func (q *queue[T]) Dequeue() (T, error) {
	var r T

	if len(q.queue) == 0 {
		return r, errors.New("empty queue")
	}

	r = q.queue[0]
	q.queue = q.queue[1:]

	return r, nil
}

func (q *queue[T]) Size() int {
	return len(q.queue)
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{queue: []T{}}
}
