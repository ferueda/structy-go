package binarytree

type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}
