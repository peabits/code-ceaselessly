package sort_test

import (
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/sort"
)

func TestRandomSequenceFrom0(t *testing.T) {
	const N = 10
	sequence := sort.RandomSequenceFrom0[int](N)
	array := [N]int{}
	for i := range array {
		array[i] = i
	}
	t.Logf("\n%v\n", sequence)
}

func TestRandomSequenceRepeatedFrom0(t *testing.T) {
	const N = 10
	sequence := sort.RandomSequenceRepeatedFrom0[int](N)
	array := [N]int{}
	for i := range array {
		array[i] = i
	}
	t.Logf("\n%v\n", sequence)
}
