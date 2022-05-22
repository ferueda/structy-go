package sumList

import (
	"github.com/ferueda/structy-go/linkedlist"
)

func sumList(node *linkedlist.Node[int]) int {
	s := 0

	for n := node; n != nil; n = n.Next {
		s += n.Val
	}

	return s
}
