package dev_test

import (
	"fmt"
	Sort "sort"
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/sort"
	sort_dev "github.com/peabits/code-ceaselessly/algorithm/sort/dev"
)

func TestRadixSort(t *testing.T) {
	list := sort.RandomSequenceRepeated[int](10, 20)
	fmt.Println(list)
	sort_dev.RadixSort(list)
	fmt.Println(list)
}

func TestRadixSort1(t *testing.T) {
	list := sort.RandomSequenceRepeated[int](10, 30)
	fmt.Println(list)
	sort_dev.RadixSort1(list)
	fmt.Println(list)
}
func TestRadixSort2(t *testing.T) {
	list := sort.RandomSequenceRepeated[int](10, 300)
	fmt.Println(list)
	sort_dev.RadixSort2(list)
	fmt.Println(list)
}
func TestList(t *testing.T) {
	source := sort.RandomSequenceRepeated[int](10, 30)
	list1 := sort_dev.New(source)
	list2 := sort_dev.New(source)
	fmt.Println("list2:", list2)
	Sort.Sort(list1)
	fmt.Println("list2:", list2)
	fmt.Println("list1:", list1)
	fmt.Println("length", list1.Len())
	fmt.Println("list", list1.List())
	fmt.Println("less", list1.Less(16, 19))
	fmt.Println("less", list1.Less(19, 16))
	list1.Swap(16, 19)
	fmt.Println("less", list1.Less(16, 19))
	fmt.Println("list", list1.List())
	list1.Append(100)
	fmt.Println("list", list1.List())
	list1.Insert(0, 100)
	fmt.Println("list", list1.List())
	list1.Clear()
	fmt.Println("list", list1.List())
	list1 = sort_dev.New(source)
	// list1.Update(10, 100)
	(*list1).Update(10, 100)
	fmt.Println("list", list1.Value(10))
	fmt.Println("list", list1.List())
}
