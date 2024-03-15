package linkedlist

import (
	"fmt"
	"strings"

	"github.com/peabits/code-ceaselessly/common"
)

type LinkedList[T any] struct {
	root   *Node[T]
	length int
}

func New[T any](list []T) *LinkedList[T] {
	var root *Node[T]
	switch len(list) {
	case 0:
		root = nil
	case 1:
		root = &Node[T]{list[0], nil}
	default:
		root = &Node[T]{list[0], nil}
		cursor := root
		for i := 1; i < len(list); i++ {
			cursor.next = &Node[T]{list[i], nil}
			cursor = cursor.next
		}
	}
	return &LinkedList[T]{
		root, len(list),
	}
}

func (llist *LinkedList[T]) List() []T {
	list := make([]T, llist.length)
	for i, cursor := 0, llist.root; i < llist.length; i, cursor = i+1, cursor.next {
		list[i] = cursor.data
	}
	return list
}

func (llist *LinkedList[T]) Interval(l, r int) []T {
	if l >= r || l < 0 || r > llist.length {
		panic("IndexError")
	}
	list := make([]T, r-l)
	cursor := llist.root
	for i := 0; i < l; i++ {
		cursor = cursor.next
	}
	for i := l; i < r; i++ {
		list[i-l] = cursor.data
		cursor = cursor.next
	}
	return list
}

func (llist *LinkedList[T]) Value(i int) T {
	if i < 0 || i >= llist.length {
		panic(common.IndexError("IndexError"))
	}
	cursor := llist.root
	for j := 0; j < i; j++ {
		cursor = cursor.next
	}
	return cursor.data
}

func (llist *LinkedList[T]) Update(i int, value T) {
	if i < 0 || i >= llist.length {
		panic(common.IndexError("IndexError"))
	}
	cursor := llist.root
	for j := 0; j < i; j++ {
		cursor = cursor.next
	}
	cursor.data = value
}

func (llist *LinkedList[T]) Swap(i, j int) {
	if i < 0 || j < 0 || i >= llist.length || j >= llist.length {
		panic("IndexError")
	}
	if i == j {
		return
	}
	if i > j {
		i, j = j, i
	}
	cursor := llist.root
	for k := 0; k < i; k++ {
		cursor = cursor.next
	}
	nodei := cursor
	for k := i; k < j; k++ {
		cursor = cursor.next
	}
	nodej := cursor
	nodei.data, nodej.data = nodej.data, nodei.data
}

func (llist *LinkedList[T]) Entry(val T) {
	cursor := llist.root
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = &Node[T]{val, nil}
	llist.length++
}

func (llist *LinkedList[T]) Insert(loc int, val T) {
	if loc < 0 || loc > llist.length+1 {
		panic(common.IndexError("Index Error"))
	}
	if loc == 0 {
		cursor := &Node[T]{val, nil}
		cursor.next = llist.root
		llist.root = cursor
	} else {
		cursor := llist.root
		for i := 1; i < loc; i++ {
			cursor = cursor.next
		}
		temp := &Node[T]{val, nil}
		temp.next = cursor.next.next
		cursor.next = temp
	}
	llist.length++
}

func (llist *LinkedList[T]) Len() int {
	return llist.length
}

func (llist *LinkedList[T]) String() string {
	builder := new(strings.Builder)
	builder.WriteString("linkedlist.LinkedList[ ")
	cursor := llist.root
	for i := 0; i < llist.length; i++ {
		builder.WriteString(fmt.Sprintf("%v", cursor.data))
		if i != llist.length-1 {
			builder.WriteString(" -> ")
		}
		cursor = cursor.next
	}
	builder.WriteString(" ]")
	return builder.String()
}
