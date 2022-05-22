package linkedListValues

import (
	"github.com/ferueda/structy-go/linkedlist"
)

func linkedListValues[T any](node *linkedlist.Node[T]) []T {
	r := []T{}

	for n := node; n != nil; n = n.Next {
		r = append(r, n.Val)
	}

	return r
}
