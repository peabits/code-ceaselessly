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
