package sort

import (
	"cmp"
)

/*
按照由小到大的顺序对一个数组的元素进行排序（使用 Go 内置排序算法 pdqsort）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, })
	fmt.Println(list) // [5 3 4 2 1]
	list = BuiltinSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func BuiltinSort[T cmp.Ordered](list ArrayList[T]) ArrayList[T] {
	list.Sort()
	return list
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用选择排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = SelectionSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func SelectionSort[T cmp.Ordered](list ArrayList[T]) ArrayList[T] {
	for i := 0; i < list.length-1; i++ {
		minIndex := i
		for j := i + 1; j < list.length; j++ {
			if list.Less(j, minIndex) {
				minIndex = j
			}
		}
		list.Swap(i, minIndex)

	}
	return list
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用冒泡排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = BubbleSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func BubbleSort[T cmp.Ordered](list ArrayList[T]) ArrayList[T] {
	for count, ordered, i := list.length-1, false, 0; !ordered && count > 0; count-- {
		for i, ordered = 0, true; i < count; i++ {
			if list.Greater(i, i+1) {
				ordered = false
				list.Swap(i, i+1)
			}
		}
	}
	return list
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用插入排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = InsertionSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func InsertionSort[T cmp.Ordered](list ArrayList[T]) ArrayList[T] {
	for i, j := 1, 0; i < list.length; i++ {
		minValue := list.Value(i)
		for j = i; j > 0 && list.Value(j-1) > minValue; j-- {
			list.Update(j, list.Value(j-1))
		}
		list.Update(j, minValue)
	}
	return list
}

/*
按照由小到大的顺序对一个数组区间 [start, end) 内的元素进行排序（使用插入排序算法）。

类型：

	func (ArrayList[T], int, int) ArrayList[T]

参数：

  - list: 等待排序的数组
  - start: 开始索引（包括）
  - end: 结束索引（不包括）

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{1, 4, 2, 3, 5})
	fmt.Println(list) // [1 4 2 3 5]
	list = insertionSortInterval(list, 1, 4)
	fmt.Println(list) // [1 2 3 4 5]
*/
func insertionSortInterval[T cmp.Ordered](list ArrayList[T], start, end int) ArrayList[T] {
	for i, j := start+1, 0; i < end; i++ {
		for j = i; j > 0 && list.Less(j, j-1); j-- {
			list.Swap(j, j-1)
		}
	}
	return list
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用计数排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = CountingSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func CountingSort[T Integer](list ArrayList[T]) ArrayList[T] {
	maxValue, minValue := list.maxValue, list.minValue
	backup := list.Copy()
	buckets := make([]int, maxValue-minValue+1)
	for i := 0; i < backup.length; i++ {
		buckets[backup.Value(i)-minValue]++
	}
	for i := 0; i < len(buckets)-1; i++ {
		buckets[i+1] += buckets[i]
	}
	for i := backup.length - 1; i >= 0; i-- {
		buckets[backup.Value(i)-minValue]--
		list.Update(buckets[backup.Value(i)-minValue], backup.Value(i))
	}
	return list
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用简化的计数排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = CountingSortShallow(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func CountingSortShallow[T Integer](list ArrayList[T]) ArrayList[T] {
	minValue, maxValue := list.minValue, list.maxValue
	buckets := make([]T, maxValue-minValue+1)
	for i := 0; i < list.length; i++ {
		buckets[list.Value(i)-minValue]++
	}
	for i, k := 0, 0; i < len(buckets); i++ {
		for j := buckets[i]; j > 0; j-- {
			list.Update(k, T(i)+minValue)
			k++
		}
	}
	return list
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用基数排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = RadixSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func RadixSort[T cmp.Ordered](list ArrayList[T]) ArrayList[T] {
	panic("Pending implementation")
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用普通的快速排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = QuickSortPlain(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func QuickSortPlain[T cmp.Ordered](list ArrayList[T]) ArrayList[T] {
	quickSortPlain(list, 0, list.length)
	return list
}

func quickSortPlain[T cmp.Ordered](list ArrayList[T], start, end int) {
	if end-start <= 1 {
		return
	}
	pivotIndex := start + random.Intn(end-start) // random index
	pivot := list.Value(pivotIndex)
	list.Swap(pivotIndex, start)
	leftIndex, rightIndex := start+1, end-1
	for leftIndex <= rightIndex {
		if list.Value(leftIndex) <= pivot {
			leftIndex++
		} else {
			list.Swap(leftIndex, rightIndex)
			rightIndex--
		}
	}
	list.Swap(start, leftIndex-1)
	quickSortPlain(list, start, leftIndex-1)
	quickSortPlain(list, rightIndex+1, end)
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用优化的快速排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = QuickSortPlain(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func QuickSortOptimized[T cmp.Ordered](list ArrayList[T]) ArrayList[T] {
	quickSortOptimized[T](list, 0, list.length)
	return list
}

func quickSortOptimized[T cmp.Ordered](list ArrayList[T], start, end int) {
	if end-start <= 3 {
		insertionSortInterval(list, start, end)
	} else {
		pivot := max(list.Value(start), list.Value(end-1), list.Value(start+(end-start)/2))
		midIndex, leftIndex, rightIndex := start, start, end-1
		for midIndex <= rightIndex {
			if v := list.Value(midIndex); v < pivot {
				list.Swap(midIndex, leftIndex)
				midIndex++
				leftIndex++
			} else if v > pivot {
				list.Swap(midIndex, rightIndex)
				rightIndex--
			} else {
				midIndex++
			}
		}
		quickSortOptimized(list, start, leftIndex)
		quickSortOptimized(list, rightIndex+1, end)
	}

}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用归并排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = MergeSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func MergeSort[T cmp.Ordered](list ArrayList[T], start, end int) {
	// panic("Pending implementation")

}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用堆排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = HeapSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func HeapSort[T cmp.Ordered](list ArrayList[T], start, end int) {
	panic("Pending implementation")
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用桶排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = BucketSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func BucketSort[T cmp.Ordered](list ArrayList[T], start, end int) {
	panic("Pending implementation")
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用希尔排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = ShellSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func ShellSort[T cmp.Ordered](list ArrayList[T], start, end int) {
	panic("Pending implementation")
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用锦标赛排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = TournamentSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func TournamentSort[T cmp.Ordered](list ArrayList[T], start, end int) {
	panic("Pending implementation")
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用 tim 排序算法）。

类型：

	func (ArrayList[T]) ArrayList[T]

参数：

  - list: 等待排序的数组

返回值:

  - 完成排序的数组

例子:

	list := NewArrayList([]int{5, 3, 4, 2, 1})
	fmt.Println(list) // [5 3 4 2 1]
	list = TimSort(list)
	fmt.Println(list) // [1 2 3 4 5]
*/
func TimSort[T cmp.Ordered](list ArrayList[T], start, end int) {
	panic("Pending implementation")
}
