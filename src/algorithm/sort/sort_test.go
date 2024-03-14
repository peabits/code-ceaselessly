package sort_test

import (
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/sort"
)

func TestBuiltinSort(t *testing.T) {
	testSortFunc(t, sort.BuiltinSort)
}

func BenchmarkBuiltinSort(b *testing.B) {
	benchmarkSortFunc(b, sort.BuiltinSort)
}

func TestSelectionSort(t *testing.T) {
	testSortFunc(t, sort.SelectionSort)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkSortFunc(b, sort.SelectionSort)
}

func TestBubbleSort(t *testing.T) {
	testSortFunc(t, sort.BubbleSort)
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmarkSortFunc(b, sort.BubbleSort)
}

func TestInsertionSort(t *testing.T) {
	testSortFunc(t, sort.InsertionSort)
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkSortFunc(b, sort.InsertionSort)
}

func TestCountingSort(t *testing.T) {
	testSortFunc(t, sort.CountingSort)
}

func BenchmarkCountingSort(b *testing.B) {
	benchmarkSortFunc(b, sort.CountingSort)
}

func TestCountingSortShallow(t *testing.T) {
	testSortFunc(t, sort.CountingSortShallow)
}

func BenchmarkCountingSortShallow(b *testing.B) {
	benchmarkSortFunc(b, sort.CountingSortShallow)
}

func TestRadixSort(t *testing.T) {
	// testSortFunc(t, sort.RadixSort)
}

func BenchmarkRadixSort(b *testing.B) {
	// benchmarkSortFunc(b, sort.RadixSort)
}

func TestQuickSortPlain(t *testing.T) {
	testSortFunc(t, sort.QuickSortPlain)
}

func BenchmarkQuickSortPlain(b *testing.B) {
	benchmarkSortFunc(b, sort.QuickSortPlain)
}

func TestQuickSortOptimized(t *testing.T) {
	testSortFunc(t, sort.QuickSortOptimized)
}

func BenchmarkQuickSortOptimized(b *testing.B) {
	benchmarkSortFunc(b, sort.QuickSortOptimized)
}

func TestMergeSort(t *testing.T) {
	testSortFunc(t, sort.MergeSort)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkSortFunc(b, sort.MergeSort)
}

func TestHeapSort(t *testing.T) {
	// testSortFunc(t, sort.HeapSort)
}

func BenchmarkHeapSort(b *testing.B) {
	// benchmarkSortFunc(b, sort.HeapSort)
}

func TestBucketSort(t *testing.T) {
	// testSortFunc(t, sort.BucketSort)
}

func BenchmarkBucketSort(b *testing.B) {
	// benchmarkSortFunc(b, sort.BucketSort)
}
func TestShellSort(t *testing.T) {
	// testSortFunc(t, sort.ShellSort)
}

func BenchmarkShellSort(b *testing.B) {
	// benchmarkSortFunc(b, sort.ShellSort)
}
func TestTournamentSort(t *testing.T) {
	// testSortFunc(t, sort.TournamentSort)
}

func BenchmarkTournamentSort(b *testing.B) {
	// benchmarkSortFunc(b, sort.TournamentSort)
}

func TestTimSort(t *testing.T) {
	// testSortFunc(t, sort.TimSort)
}

func BenchmarkTimSort(b *testing.B) {
	// benchmarkSortFunc(b, sort.TimSort)
}
