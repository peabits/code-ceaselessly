package dev

import "cmp"

func ShellSort0[T cmp.Ordered](list []T) {
	gap := 1
	for g := gap<<1 + 1; g < len(list); g = g<<1 + 1 { // 2^k+1
		gap = g
	}
	for ; gap > 0; gap = (gap - 1) >> 1 {
		for i := gap; i < len(list); i++ {
			for j := i; j >= gap && list[j] < list[j-gap]; j -= gap {
				list[j], list[j-gap] = list[j-gap], list[j]
			}
		}
	}
}

func ShellSort1[T cmp.Ordered](list []T) {
	gap := 1
	for g := gap<<1 + 1; g < len(list); g = g<<1 + 1 { // 2^k+1
		gap = g
	}
	for ; gap > 0; gap = (gap - 1) >> 1 {
		for k := 0; k < gap; k++ {
			for i := k + gap; i < len(list); i += gap {
				for j := i; j > k && list[j] < list[j-gap]; j -= gap {
					list[j], list[j-gap] = list[j-gap], list[j]
				}
			}
		}
	}
}
