package dp_test

import (
	"testing"

	"github.com/peabits/code-ceaselessly/algorithm/dp"
)

func TestDigitalTrianglesProblem(t *testing.T) {
	cases := [][]int{
		{7, 3, 8, 8, 1, 0, 2, 7, 4, 4, 4, 5, 2, 6, 5},
	}
	expect := []int{30}
	for i, input := range cases {
		output := dp.DigitalTrianglesProblemDP(input)
		if output == expect[i] {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Sucess.\n", i, input, output, expect[i])
		} else {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Fail.\n", i, input, output, expect[i])
		}
	}
}

func TestGatherHerbsProblemDFS(t *testing.T) {
	cases := [][][]int{
		{{70, 3}, {71, 100}, {69, 1}, {1, 2}},
	}
	expect := []int{3}
	for i, input := range cases {
		output := dp.GatherHerbsProblemDFS(input)
		if output == expect[i] {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Sucess.\n", i, input, output, expect[i])
		} else {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Fail.\n", i, input, output, expect[i])
		}
	}
}

func TestGatherHerbsProblemMS(t *testing.T) {
	cases := [][][]int{
		{{70, 3}, {71, 100}, {69, 1}, {1, 2}},
	}
	expect := []int{3}
	for i, input := range cases {
		output := dp.GatherHerbsProblemMS(input)
		if output == expect[i] {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Sucess.\n", i, input, output, expect[i])
		} else {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Fail.\n", i, input, output, expect[i])
		}
	}
}

func TestGatherHerbsProblemDP(t *testing.T) {
	cases := [][][]int{
		{{70, 3}, {71, 100}, {69, 1}, {1, 2}},
	}
	expect := []int{3}
	for i, input := range cases {
		output := dp.GatherHerbsProblemDP(input)
		if output == expect[i] {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Sucess.\n", i, input, output, expect[i])
		} else {
			t.Logf("\nCase%2d: %v\nOutput: %v\nExpect: %v\nResult: Fail.\n", i, input, output, expect[i])
		}
	}
}
