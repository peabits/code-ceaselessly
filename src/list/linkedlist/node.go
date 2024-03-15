package linkedlist

type Node[T any] struct {
	data T
	next *Node[T]
}

// func createNode[T any](data T, next *Node[T]) *Node[T] {
// 	return &Node[T]{
// 		data, next,
// 	}
// }
