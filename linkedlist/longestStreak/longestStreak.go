package longestStreak

import (
	"github.com/ferueda/structy-go/linkedlist"
)

func longestStreak[T comparable](node *linkedlist.Node[T]) int {
	current, longest := 0, 0
	var prev T
	for n := node; n != nil; n = n.Next {
		if n.Val == prev {
			current += 1
		} else {
			current = 1
		}

		if current > longest {
			longest = current
		}

		prev = n.Val
	}

	return longest
}
