package sort_test

import (
	"math/rand"
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/sort"
)

func testSortFunc(t *testing.T, sortFunc func(list sort.ArrayList[int])) {
	var N = (rand.Intn(10) + 1) * 100
	sequence := sort.RandomSequenceRepeatedFrom0[int](N)
	listSource := sort.NewArrayList(sequence)
	listSorted := listSource.Copy()
	listSource.Sort()
	listSorted.SortFunc(sortFunc)
	if listSorted.EqualTo(listSource) {
		t.Logf("\nSource: %v\nSorted: %v\nResult: Success\n", listSource, listSorted)
	} else {
		t.Fatalf("\nSource: %v\nSorted: %v\nResult: Error\n", listSource, listSorted)
	}
}

func benchmarkSortFunc(b *testing.B, sortFunc func(list sort.ArrayList[int])) {
	var N = 10000
	sequence := sort.RandomSequenceRepeatedFrom0[int](N)
	list := sort.NewArrayList(sequence)
	for i := 0; i < b.N; i++ {
		sortFunc(list)
	}
}
