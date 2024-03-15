package sort

import (
	"cmp"
	"sort"
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
func BuiltinSort[T cmp.Ordered](list ArrayList[T]) {
	list.Sort()
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
func BubbleSort[T cmp.Ordered](list ArrayList[T]) {
	bubbleSort[T](list)
}

func bubbleSort[T cmp.Ordered](list sort.Interface) {
	bubbleSortInterval[T](list, 0, list.Len())
}

// 左闭右开区间 [a, b)
func bubbleSortInterval[T cmp.Ordered](list sort.Interface, a, b int) {
	for i, c := b-1, false; !c && i > a; i-- {
		c = true
		for j := a; j < i; j++ {
			if list.Less(j+1, j) {
				c = false
				list.Swap(j, j+1)
			}
		}
	}
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
func SelectionSort[T cmp.Ordered](list ArrayList[T]) {
	selectionSort[T](list)
}

func selectionSort[T cmp.Ordered](list sort.Interface) {
	selectionSortInterval[T](list, 0, list.Len())
}

// 左闭右开区间 [a, b)
func selectionSortInterval[T cmp.Ordered](list sort.Interface, a, b int) {
	for i := a; i < b-1; i++ {
		k := i
		for j := i + 1; j < b; j++ {
			if list.Less(j, k) {
				k = j
			}
		}
		list.Swap(i, k)
	}
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
func HeapSort[T cmp.Ordered](list ArrayList[T]) {
	heapSort[T](list)
}

func heapSort[T cmp.Ordered](list sort.Interface) {
	heapSortInterval[T](list, 0, list.Len())
}

// 左闭右开区间 [a, b)
func heapSortInterval[T cmp.Ordered](list sort.Interface, a, b int) {
	for i := ((b - a) - 2) >> 1; i >= 0; i-- {
		maxHeapify[T](list, a+i, b)
	}
	for i := b - a - 1; i >= 0; i-- {
		list.Swap(0, i)
		maxHeapify[T](list, 0, i)
	}
}

// 左闭右开区间 [a, b)
func maxHeapify[T cmp.Ordered](list sort.Interface, a, b int) {
	for m, n := a, a<<1+1; n < b; m, n = n, n<<1+1 {
		if n+1 < b && list.Less(n, n+1) {
			n++
		}
		if list.Less(n, m) {
			return
		}
		list.Swap(m, n)
	}
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
func InsertionSort[T cmp.Ordered](list ArrayList[T]) {
	insertionSort[T](list)
}

func insertionSort[T cmp.Ordered](list sort.Interface) {
	insertionSortInterval[T](list, 0, list.Len())
}

// 左闭右开区间 [a, b)
func insertionSortInterval[T cmp.Ordered](list sort.Interface, a, b int) {
	for i, j := a+1, 0; i < b; i++ {
		for j = i; j > 0 && list.Less(j, j-1); j-- {
			list.Swap(j, j-1)
		}
	}
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
func ShellSort[T cmp.Ordered](list ArrayList[T]) {
	shellSort[T](list)
}

func shellSort[T cmp.Ordered](list sort.Interface) {
	shellSortInterval[T](list, 0, list.Len())
}

func shellSortInterval[T cmp.Ordered](list sort.Interface, a, b int) {
	gap := 1
	for g := gap<<1 + 1; g < b-a; g = g<<1 + 1 {
		gap = g
	}
	for ; gap > 0; gap = (gap - 1) >> 1 {
		for i := a + gap; i < b; i++ {
			for j := i; j >= a+gap && list.Less(j, j-gap); j -= gap {
				list.Swap(j, j-gap)
			}
		}
	}
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
func QuickSortPlain[T cmp.Ordered](list ArrayList[T]) {
	quickSortPlain[T](list)
}

func quickSortPlain[T cmp.Ordered](list sort.Interface) {
	quickSortPlainInterval[T](list, 0, list.Len())
}

// 左闭右开区间 [a, b)
func quickSortPlainInterval[T cmp.Ordered](list sort.Interface, a, b int) {
	if b-a <= 1 {
		return
	}
	pi := a + random.Intn(b-a) // pivot => random index
	list.Swap(pi, a)
	l, r := a+1, b-1
	for l <= r {
		if list.Less(a, l) {
			list.Swap(l, r)
			r--
		} else {
			l++
		}
	}
	list.Swap(a, l-1)
	quickSortPlainInterval[T](list, a, l-1)
	quickSortPlainInterval[T](list, r+1, b)
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
func QuickSortOptimized[T cmp.Ordered](list ArrayList[T]) {
	quickSortOptimized[T](list)
}

func quickSortOptimized[T cmp.Ordered](list sort.Interface) {
	quickSortOptimizedInterval[T](list, 0, list.Len())
}

// 左闭右开区间 [a, b)
func quickSortOptimizedInterval[T cmp.Ordered](list sort.Interface, a, b int) {
	quickSortThreshold := 10
	if b-a <= quickSortThreshold {
		insertionSortInterval[T](list, a, b)
	} else {
		pi := a + (b-a)>>1
		if list.Less(a, b-1) {
			if list.Less(pi, a) {
				pi = a
			} else if list.Less(b-1, pi) {
				pi = b - 1
			}
		} else {
			if list.Less(pi, b-1) {
				pi = b - 1
			} else if list.Less(a, pi) {
				pi = a
			}
		}
		list.Swap(a, pi)
		m, l, r := a+1, a, b-1
		for m <= r {
			if list.Less(m, m-1) {
				list.Swap(m, l)
				m++
				l++
			} else if list.Less(m-1, m) {
				list.Swap(m, r)
				r--
			} else {
				m++
			}
		}
		quickSortOptimizedInterval[T](list, a, l)
		quickSortOptimizedInterval[T](list, r+1, b)
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
func MergeSort[T cmp.Ordered](list ArrayList[T]) {
	mergeSort(list, 0, list.Len())
}

// 左闭右开区间 [a, b)
func mergeSort[T cmp.Ordered](list ArrayList[T], a, b int) {
	if b-a <= 1 {
		return
	}
	m := a + (b-a)>>1
	mergeSort(list, a, m)
	mergeSort(list, m, b)
	temp := make([]T, b-a)
	copy(temp, list.Interval(a, b))
	i, j, k := 0, m-a, a
	for ; i < m-a && j < b-a; k++ {
		if temp[i] <= temp[j] {
			list.Update(k, temp[i])
			i++
		} else {
			list.Update(k, temp[j])
			j++
		}
	}
	for ; i < m-a; k++ {
		list.Update(k, temp[i])
		i++
	}
	for ; j < b-a; k++ {
		list.Update(k, temp[j])
		j++
	}
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
func BucketSort[T Number](alist ArrayList[T]) {
	mi, ma := alist.Min()-1, alist.Max()+1
	size := 3
	N := int(ma-mi+1)/size + 1
	buckets := make([][]T, N)
	for i := range buckets {
		buckets[i] = make([]T, 0)
	}
	for i := 0; i < alist.Len(); i++ {
		v := alist.Value(i)
		bi := int(v-mi+1) / size
		buckets[bi] = append(buckets[bi], v)
	}
	i := 0
	for _, list := range buckets {
		for i := 1; i < len(list); i++ { // 简单插入排序
			for j := i; j > 0 && list[j] < list[j-1]; j-- {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
		for _, val := range list {
			alist.Update(i, val)
			i++
		}
	}
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
func CountingSort[T Integer](list ArrayList[T]) {
	mi, ma := list.Min(), list.Max()
	backup := list.Copy()
	buckets := make([]int, ma-mi+1)
	for i := 0; i < backup.Len(); i++ {
		buckets[backup.Value(i)-mi]++
	}
	for i := 0; i < len(buckets)-1; i++ {
		buckets[i+1] += buckets[i]
	}
	for i := backup.Len() - 1; i >= 0; i-- {
		buckets[backup.Value(i)-mi]--
		list.Update(buckets[backup.Value(i)-mi], backup.Value(i))
	}
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
func CountingSortShallow[T Integer](list ArrayList[T]) {
	mi, ma := list.Min(), list.Max()
	buckets := make([]T, ma-mi+1)
	for i := 0; i < list.Len(); i++ {
		buckets[list.Value(i)-mi]++
	}
	for i, k := 0, 0; i < len(buckets); i++ {
		for j := buckets[i]; j > 0; j-- {
			list.Update(k, T(i)+mi)
			k++
		}
	}
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
func RadixSort[T Integer](list ArrayList[T]) {
	minN, maxN := int(list.Min()), int(list.Max())
	inc := 0
	if minN < 0 {
		inc = -minN
	}
	for base := 1; maxN/base > 0; base *= 10 { // 计数排序
		buckets := make([]int, 10)
		for i := 0; i < list.Len(); i++ {
			n := (int(list.Value(i)) + inc) / base % 10
			buckets[n]++
		}
		for i := 1; i < len(buckets); i++ {
			buckets[i] += buckets[i-1]
		}
		temp := list.List()
		for i := list.Len() - 1; i >= 0; i-- {
			n := (int(temp[i]) + inc) / base % 10
			buckets[n]--
			list.Update(buckets[n], temp[i])
		}
	}

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
func TournamentSort[T cmp.Ordered](list ArrayList[T]) {
	panic("Pending implementation")
}

/*
按照由小到大的顺序对一个数组的元素进行排序（使用 Tim 排序算法）。

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
func TimSort[T cmp.Ordered](list ArrayList[T]) {
	panic("Pending implementation")
}

func PdqSort[T cmp.Ordered](list ArrayList[T]) {
	panic("Pending implementation")
}
