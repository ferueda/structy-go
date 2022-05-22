package linkedlist

import (
	"fmt"
)

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) String() string {
	if n == nil {
		return "nil"
	}

	r := ""

	for n := n; n != nil; n = n.Next {
		r += fmt.Sprintf("%v -> ", n.Val)
	}

	return r + "nil"
}

func (n *Node[T]) GetValues() []T {
	values := []T{}

	for curr := n; curr != nil; curr = curr.Next {
		values = append(values, curr.Val)
	}

	return values
}
