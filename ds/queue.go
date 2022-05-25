package ds

import (
	"errors"
)

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Enqueue(el T) {
	q.queue = append(q.queue, el)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var r T

	if len(q.queue) == 0 {
		return r, errors.New("empty queue")
	}

	r = q.queue[0]
	q.queue = q.queue[1:]

	return r, nil
}

func (q *Queue[T]) Size() int {
	return len(q.queue)
}
