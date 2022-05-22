package findValue

import (
	"github.com/ferueda/structy-go/linkedlist"
)

func findValue[T comparable](node *linkedlist.Node[T], target T) bool {
	for n := node; n != nil; n = n.Next {
		if n.Val == target {
			return true
		}
	}
	return false
}
