package dp_test

import (
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/dp"
)

func TestLongestCommonSubsequenceProblemDP(t *testing.T) {
	cases := [][]string{
		{"abcde", "ace"},
		{"abc", "abc"},
		{"abc", "def"},
	}
	expect := []int{
		3, 3, 0,
	}
	for i, input := range cases {
		output := dp.LongestCommonSubsequenceProblemDP(input[0], input[1])
		if output == expect[i] {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Sucess.\n", i, input, output, expect[i])
		} else {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Fail.\n", i, input, output, expect[i])
		}
	}
}
