package sort

import (
	"cmp"
	"slices"
	"sort"

	"github.com/peabits/code-ceaselessly/list/arraylist"
	"github.com/peabits/code-ceaselessly/list/linkedlist"
)

type ArrayList[T cmp.Ordered] struct {
	values *arraylist.ArrayList[T]
}

func NewArrayList[T cmp.Ordered](source []T) ArrayList[T] {
	return ArrayList[T]{arraylist.New(source)}
}

func (alist ArrayList[T]) IsNull() bool {
	return alist.values == nil
}

func (alist ArrayList[T]) Len() int {
	return alist.values.Len()
}

func (alist ArrayList[T]) List() []T {
	return alist.values.List()
}

func (alist ArrayList[T]) Interval(i, j int) []T {
	return alist.values.Interval(i, j)
}

func (alist ArrayList[T]) Value(i int) T {
	return (*alist.values)[i]
}

func (alist ArrayList[T]) Update(i int, value T) {
	(*alist.values)[i] = value
}

func (alist ArrayList[T]) Less(i, j int) bool {
	return (*alist.values)[i] < (*alist.values)[j]
}

func (alist ArrayList[T]) LessEq(i, j int) bool {
	return (*alist.values)[i] <= (*alist.values)[j]
}

func (alist ArrayList[T]) Greater(i, j int) bool {
	return (*alist.values)[i] > (*alist.values)[j]
}

func (alist ArrayList[T]) GreaterEq(i, j int) bool {
	return (*alist.values)[i] >= (*alist.values)[j]
}

func (alist ArrayList[T]) Equal(i, j int) bool {
	return (*alist.values)[i] == (*alist.values)[j]
}

func (alist ArrayList[T]) EqualTo(other ArrayList[T]) bool {
	if len(*alist.values) != len(*other.values) {
		return false
	}
	for i := 0; i < len(*alist.values); i++ {
		if (*alist.values)[i] != (*other.values)[i] {
			return false
		}
	}
	return true
}

func (alist ArrayList[T]) Swap(i, j int) {
	alist.values.Swap(i, j)
}

func (alist ArrayList[T]) Sort() {
	sort.Sort(alist) // Or slices.Sort(*alist.values)
}

func (alist ArrayList[T]) SortFunc(sortFunc func(ArrayList[T])) {
	if sortFunc == nil {
		alist.Sort()
	} else {
		sortFunc(alist)
	}
}

func (alist ArrayList[T]) Copy() ArrayList[T] {
	return ArrayList[T]{(*alist.values).Copy()}
}

func (alist ArrayList[T]) Min() T {
	return slices.Min(*alist.values)
}

func (alist ArrayList[T]) MinAb(a, b int) T {
	return slices.Min((*alist.values)[a:b])
}

func (alist ArrayList[T]) Max() T {
	return slices.Max(*alist.values)
}

func (alist ArrayList[T]) MaxAb(a, b int) T {
	return slices.Max((*alist.values)[a:b])
}

func (alist ArrayList[T]) String() string {
	return alist.values.String()
}

func (alist ArrayList[T]) Entry(val T) {
	alist.values.Entry(val)
}

type LinkedList[T cmp.Ordered] struct {
	values *linkedlist.LinkedList[T]
}

func NewLinkedList[T cmp.Ordered](source []T) LinkedList[T] {
	return LinkedList[T]{linkedlist.New(source)}
}

func (llist LinkedList[T]) Len() int {
	return llist.values.Len()
}

func (llist LinkedList[T]) List() []T {
	return llist.values.List()
}

func (llist LinkedList[T]) Interval(i, j int) []T {
	return llist.values.Interval(i, j)
}

func (llist LinkedList[T]) Less(i, j int) bool {
	return llist.values.Value(i) < llist.values.Value(j)
}

func (llist LinkedList[T]) LessEq(i, j int) bool {
	return llist.values.Value(i) <= llist.values.Value(j)
}

func (llist LinkedList[T]) Greater(i, j int) bool {
	return llist.values.Value(i) > llist.values.Value(j)
}

func (llist LinkedList[T]) GreaterEq(i, j int) bool {
	return llist.values.Value(i) >= llist.values.Value(j)
}

func (llist LinkedList[T]) Equal(i, j int) bool {
	return llist.values.Value(i) == llist.values.Value(j)
}

func (llist LinkedList[T]) Entry(v T) {
	llist.values.Entry(v)
}

func (llist LinkedList[T]) Swap(i, j int) {
	llist.values.Swap(i, j)
}

func (llist LinkedList[T]) String() string {
	return llist.values.String()
}
