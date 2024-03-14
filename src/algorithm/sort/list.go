package sort

import (
	"cmp"
	"fmt"
	"slices"
	"sort"

	"github.com/peabits/code-ceaselessly/arraylist"
)

type List[T cmp.Ordered] interface {
	Len() int
	At(int) T
	Update(int, T) bool
	Insert(int, T) bool
	Delete(int) (T, error)
	Sort()
	Push(T)
	Pop()
}

type ArrayList[T cmp.Ordered] struct {
	arraylist.ArrayList[T]
	maxValue, minValue T
	length             int
}

func NewArrayList[T cmp.Ordered](values []T) ArrayList[T] {
	return ArrayList[T]{arraylist.New(values), slices.Max(values), slices.Min(values), len(values)}
}
func (list ArrayList[T]) Less(i, j int) bool {
	return list.Value(i) < list.Value(j)
}

func (list ArrayList[T]) Greater(i, j int) bool {
	return list.Value(i) > list.Value(j)
}

func (list ArrayList[T]) Equal(i, j int) bool {
	return list.Value(i) == list.Value(j)
}

func (list ArrayList[T]) EqualTo(other ArrayList[T]) bool {
	if list.length != other.length {
		return false
	}
	for i := 0; i < list.length-1; i++ {
		if list.Value(i) != other.Value(i) {
			return false
		}
	}
	return true
}

func (list ArrayList[T]) Sort() {
	sort.Sort(list)
	// slices.Sort(list.List())
}

func (list ArrayList[T]) SortFunc(sortFunc func(ArrayList[T]) ArrayList[T]) {
	if sortFunc == nil {
		list.Sort()
	} else {
		_ = sortFunc(list)
	}
}

func (list ArrayList[T]) Copy() ArrayList[T] {
	return ArrayList[T]{list.ArrayList.Copy(), list.maxValue, list.minValue, list.length}
}

func (list ArrayList[T]) MaxValue() T {
	return list.maxValue
}

func (list ArrayList[T]) MinValue() T {
	return list.minValue
}

func (list ArrayList[T]) String() string {
	return fmt.Sprintf("%v", list.ArrayList)
}
