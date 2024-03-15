package arraylist

import (
	"bytes"
	"fmt"
)

type ArrayList[T any] []T

func New[T any](source []T) *ArrayList[T] {
	values := make([]T, len(source))
	_ = copy(values, source)
	alist := ArrayList[T](values)
	return &alist
}

func (alist ArrayList[T]) Len() int {
	return len(alist)
}

func (alist ArrayList[T]) List() []T {
	listCopy := make([]T, len(alist))
	_ = copy(listCopy, alist)
	return listCopy
}

func (alist ArrayList[T]) Interval(l, r int) []T {
	if l >= r || l < 0 || r > len(alist) {
		panic("Error: start >= end Or start < 0")
	}
	listCopy := make([]T, r-l)
	_ = copy(listCopy, alist[l:r])
	return listCopy
}

func (alist ArrayList[T]) Value(i int) T {
	return alist[i]
}

func (alist ArrayList[T]) First(i int) T {
	return alist[0]
}

func (alist ArrayList[T]) Last(i int) T {
	return alist[len(alist)-1]
}

func (alist ArrayList[T]) Update(i int, val T) {
	alist[i] = val
}

func (alist ArrayList[T]) Swap(i, j int) {
	alist[i], alist[j] = alist[j], alist[i]
}

func (alist ArrayList[T]) Copy() *ArrayList[T] {
	listCopy := make([]T, len(alist))
	_ = copy(listCopy, alist)
	alist1 := ArrayList[T](listCopy)
	return &alist1
}

func (alist ArrayList[T]) String() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("arraylist.ArrayList[ ")
	for i, v := range alist {
		buffer.WriteString(fmt.Sprintf("%v", v))
		if i != len(alist)-1 {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString(" ]")
	return buffer.String()
}

func (alist *ArrayList[T]) Entry(val T) {
	*alist = append(*alist, val)
}
