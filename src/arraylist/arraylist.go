package arraylist

import "fmt"

type ArrayList[T any] struct {
	values []T
}

func New[T any](values []T) ArrayList[T] {
	return ArrayList[T]{values}
}

func (list ArrayList[T]) Len() int {
	return len(list.values)
}

func (list ArrayList[T]) List() []T {
	listCopy := make([]T, list.Len())
	copy(listCopy, list.values)
	return listCopy
}

func (list ArrayList[T]) at(i int) *T {
	return &list.values[i]
}

func (list ArrayList[T]) Value(i int) T {
	return *list.at(i)
}

func (list ArrayList[T]) Update(i int, value T) {
	list.values[i] = value
}

func (list ArrayList[T]) Swap(i, j int) {
	list.values[i], list.values[j] = list.values[j], list.values[i]
}

func (list ArrayList[T]) Copy() ArrayList[T] {
	return ArrayList[T]{list.List()}
}

func (list ArrayList[T]) String() string {
	return fmt.Sprintf("List%v", list.values)
}
