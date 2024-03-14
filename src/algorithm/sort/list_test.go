package sort_test

import (
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/sort"
)

func TestArrayListLess(t *testing.T) {
	list := sort.NewArrayList([]int{1, 2, 1})
	if list.Less(0, 1) {
		t.Logf("\nSuccess\n")
	} else {
		t.Fatalf("\nError\n")
	}
}

func TestArrayListGreater(t *testing.T) {
	list := sort.NewArrayList([]int{1, 2, 1})
	if list.Greater(1, 2) {
		t.Logf("\nSuccess\n")
	} else {
		t.Fatalf("\nError\n")
	}
}

func TestArrayListEqual(t *testing.T) {
	list := sort.NewArrayList([]int{1, 2, 1})
	if list.Equal(0, 2) {
		t.Logf("\nSuccess\n")
	} else {
		t.Fatalf("\nError\n")
	}
}
