package dev

import (
	"cmp"
	"slices"
)

type List []int

func New(source []int) *List {
	N := 10
	if len(source) > N {
		N = len(source)
	}
	list := List(make([]int, N))
	copy(list, source)
	return &list
}

func (list List) Len() int {
	return len(list)
}

func (list List) Less(i, j int) bool {
	return list[i] < list[j]
}

func (list List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list *List) Append(v int) {
	*list = append(*list, v)
}

func (list *List) Insert(i int, v int) {
	*list = append(*list, v)
	copy((*list)[i+1:], (*list)[i:])
	(*list)[i] = v
}

func (list *List) Clear() {
	*list = make(List, 0, 10)
}

func (list *List) Update(i int, v int) {
	(*list)[i] = v
}

func (list List) Value(i int) int {
	return list[i]
}

func (list List) List() []int {
	list1 := make([]int, len(list))
	copy(list1, list)
	return list1
}

func ShellSort0[T cmp.Ordered](list []T) {
	gap := 1
	for g := gap<<1 + 1; g < len(list); g = g<<1 + 1 { // 2^k+1
		gap = g
	}
	for ; gap > 0; gap = (gap - 1) >> 1 {
		for i := gap; i < len(list); i++ {
			for j := i; j >= gap && list[j] < list[j-gap]; j -= gap {
				list[j], list[j-gap] = list[j-gap], list[j]
			}
		}
	}
}

func ShellSort1[T cmp.Ordered](list []T) {
	gap := 1
	for g := gap<<1 + 1; g < len(list); g = g<<1 + 1 { // 2^k+1
		gap = g
	}
	for ; gap > 0; gap = (gap - 1) >> 1 {
		for k := 0; k < gap; k++ {
			for i := k + gap; i < len(list); i += gap {
				for j := i; j > k && list[j] < list[j-gap]; j -= gap {
					list[j], list[j-gap] = list[j-gap], list[j]
				}
			}
		}
	}
}

func bitSort(arr []int, bit int) []int {
	n := len(arr)

	bitCounts := make([]int, 10)
	for i := 0; i < n; i++ {
		pos := (arr[i] / bit) % 10
		bitCounts[pos]++
	}
	for i := 0; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1]
	}

	tmp := make([]int, 10)
	for i := n - 1; 0 <= i; i-- {
		pos := (arr[i] / bit) % 10
		tmp[bitCounts[pos]-1] = arr[i]
		bitCounts[pos]--
	}
	for i := 0; i < n; i++ {
		arr[i] = tmp[i]
	}

	return arr
}

// 基数排序
func RadixSort(arr []int) []int {
	max := slices.Max(arr)
	for bit := 1; max/bit > 0; bit *= 10 {
		arr = bitSort(arr, bit)
	}
	return arr
}

func RadixSort1(list []int) {
	max := slices.Max(list)
	for digit := 1; max/digit > 0; digit *= 10 {
		buckets := make([][]int, 10)
		for i := range buckets {
			buckets[i] = make([]int, 0, 10)
		}
		for i := 0; i < len(list); i++ {
			n := list[i] / digit % 10
			buckets[n] = append(buckets[n], list[i])
		}
		i := 0
		for _, bucket := range buckets {
			for _, el := range bucket {
				list[i] = el
				i++
			}
		}
	}
}

func RadixSort2(list List) {
	max := slices.Max(list)
	for digit := 1; max/digit > 0; digit *= 10 {
		buckets := make([]int, 10)
		for i := 0; i < len(list); i++ {
			n := list[i] / digit % 10
			buckets[n]++
		}
		for i := 1; i < len(buckets); i++ {
			buckets[i] += buckets[i-1]
		}
		temp := list.List()
		for i := len(temp) - 1; i >= 0; i-- {
			n := temp[i] / digit % 10
			buckets[n]--
			list.Update(buckets[n], temp[i])
		}
	}
}
