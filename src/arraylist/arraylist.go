package arraylist

import "fmt"

type ArrayList[T any] struct {
	values []T
	length int
}

func New[T any](values []T) ArrayList[T] {
	return ArrayList[T]{values, len(values)}
}

func (list ArrayList[T]) Len() int {
	return list.length
}

func (list ArrayList[T]) List() []T {
	listCopy := make([]T, list.length)
	_ = copy(listCopy, list.values)
	return listCopy
}

func (list ArrayList[T]) Interval(l, r int) []T {
	if l >= r || l < 0 {
		panic("Error: start >= end")
	}
	listCopy := make([]T, r-l)
	_ = copy(listCopy, list.values[l:r])
	return listCopy
}

func (list ArrayList[T]) at(i int) *T {
	return &list.values[i]
}

func (list ArrayList[T]) Value(i int) T {
	return *list.at(i)
}

func (list ArrayList[T]) First(i int) T {
	return list.values[0]
}

func (list ArrayList[T]) Last(i int) T {
	return list.values[list.length-1]
}

func (list ArrayList[T]) Update(i int, value T) {
	list.values[i] = value
}

func (list ArrayList[T]) Swap(i, j int) {
	list.values[i], list.values[j] = list.values[j], list.values[i]
}

func (list ArrayList[T]) Copy() ArrayList[T] {
	return ArrayList[T]{list.List(), list.length}
}

func (list ArrayList[T]) String() string {
	return fmt.Sprintf("List%v", list.values)
}
