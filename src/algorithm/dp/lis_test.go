package dp_test

import (
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/dp"
)

func TestLengthOfLISProblemDP(t *testing.T) {
	cases := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
	}
	expect := []int{4, 4, 1}
	for i, input := range cases {
		output := dp.LengthOfLISProblemDP(input)
		if output == expect[i] {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Sucess.\n", i, input, output, expect[i])
		} else {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Fail.\n", i, input, output, expect[i])
		}
	}
}

func TestLengthOfLISProblemGA(t *testing.T) {
	cases := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
	}
	expect := []int{4, 4, 1}
	for i, input := range cases {
		output := dp.LengthOfLISProblemGA(input)
		if output == expect[i] {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Sucess.\n", i, input, output, expect[i])
		} else {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Fail.\n", i, input, output, expect[i])
		}
	}
}
