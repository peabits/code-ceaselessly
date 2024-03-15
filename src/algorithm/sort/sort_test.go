package sort_test

import (
	"math/rand"
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/sort"
)

func GoTestSortFunc(t *testing.T, sortFunc func(list sort.ArrayList[int])) {
	var N = (rand.Intn(10) + 1) * 100
	sequence := sort.RandomSequenceRepeatedFrom0[int](N)
	listSource := sort.NewArrayList(sequence)
	listSorted := listSource.Copy()
	listSource.Sort()
	listSorted.SortFunc(sortFunc)
	if listSorted.EqualTo(listSource) {
		t.Logf("\nResult: Successful\nSource: %v\nSorted: %v\n", listSource, listSorted)
	} else {
		t.Fatalf("\nResult: Failed\nSource: %v\nSorted: %v\n", listSource, listSorted)
	}
}

func GoTtestSortFuncN(t *testing.T, sortFunc func(list sort.ArrayList[int]), N int) {
	sequence := sort.RandomSequenceRepeatedFrom0[int](N)
	listSource := sort.NewArrayList(sequence)
	listSorted := listSource.Copy()
	listSource.Sort()
	listSorted.SortFunc(sortFunc)
	if listSorted.EqualTo(listSource) {
		t.Logf("\nResult: Successful\nSource: %v\nSorted: %v\n", listSource, listSorted)
	} else {
		t.Fatalf("\nResult: Failed\nSource: %v\nSorted: %v\n", listSource, listSorted)
	}
}

func TestBuiltinSort(t *testing.T) {
	GoTestSortFunc(t, sort.BuiltinSort)
}

func TestBubbleSort(t *testing.T) {
	GoTestSortFunc(t, sort.BubbleSort)
}

func TestSelectionSort(t *testing.T) {
	GoTestSortFunc(t, sort.SelectionSort)
}

func TestHeapSort(t *testing.T) {
	GoTestSortFunc(t, sort.HeapSort)
}

func TestInsertionSort(t *testing.T) {
	GoTestSortFunc(t, sort.InsertionSort)
}

func TestShellSort(t *testing.T) {
	GoTestSortFunc(t, sort.ShellSort)
}

func TestQuickSortPlain(t *testing.T) {
	GoTestSortFunc(t, sort.QuickSortPlain)
}

func TestQuickSortOptimized(t *testing.T) {
	GoTestSortFunc(t, sort.QuickSortOptimized)
}

func TestMergeSort(t *testing.T) {
	GoTestSortFunc(t, sort.MergeSort)
}

func TestBucketSort(t *testing.T) {
	GoTestSortFunc(t, sort.BucketSort)
}

func TestCountingSort(t *testing.T) {
	GoTestSortFunc(t, sort.CountingSort)
}

func TestCountingSortShallow(t *testing.T) {
	GoTestSortFunc(t, sort.CountingSortShallow)
}

func TestRadixSort(t *testing.T) {
	GoTestSortFunc(t, sort.RadixSort)
}

func TestTournamentSort(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Log(p)
		}
	}()
	GoTestSortFunc(t, sort.TournamentSort)
}

func TestTimSort(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Log(p)
		}
	}()
	GoTestSortFunc(t, sort.TimSort)
}

func TestPdqSort(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Log(p)
		}
	}()
	GoTestSortFunc(t, sort.PdqSort)
}
