package sort_test

import (
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/sort"
)

var bN int
var bList sort.ArrayList[int]

func init() {
	bN = 1000
	sequence := sort.RandomSequenceRepeatedFrom0[int](bN)
	bList = sort.NewArrayList(sequence)
}

func GoBenchmarkSortFunc(b *testing.B, sortFunc func(list sort.ArrayList[int])) {
	for i := 0; i < b.N; i++ {
		sortFunc(bList)
	}
}

func BenchmarkBuiltinSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.BuiltinSort)
}

func BenchmarkSelectionSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.SelectionSort)
}

func BenchmarkBubbleSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.BubbleSort)
}

func BenchmarkInsertionSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.InsertionSort)
}

func BenchmarkCountingSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.CountingSort)
}

func BenchmarkCountingSortShallow(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.CountingSortShallow)
}

func BenchmarkRadixSort(b *testing.B) {
	defer func() {
		if p := recover(); p != nil {
			b.Log(p)
		}
	}()
	GoBenchmarkSortFunc(b, sort.RadixSort)
}

func BenchmarkQuickSortPlain(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.QuickSortPlain)
}

func BenchmarkQuickSortOptimized(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.QuickSortOptimized)
}

func BenchmarkMergeSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.MergeSort)
}

func BenchmarkHeapSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.HeapSort)
}

func BenchmarkBucketSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.BucketSort)
}

func BenchmarkShellSort(b *testing.B) {
	GoBenchmarkSortFunc(b, sort.ShellSort)
}

func BenchmarkTournamentSort(b *testing.B) {
	defer func() {
		if p := recover(); p != nil {
			b.Log(p)
		}
	}()
	GoBenchmarkSortFunc(b, sort.TournamentSort)
}

func BenchmarkTimSort(b *testing.B) {
	defer func() {
		if p := recover(); p != nil {
			b.Log(p)
		}
	}()
	GoBenchmarkSortFunc(b, sort.TimSort)
}

func BenchmarkPdqSort(b *testing.B) {
	defer func() {
		if p := recover(); p != nil {
			b.Log(p)
		}
	}()
	GoBenchmarkSortFunc(b, sort.PdqSort)
}
