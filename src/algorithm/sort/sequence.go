package sort

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

func shuffle[T any](arr []T) []T {
	random.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func Sequence[T Integer](start, end int) []T {
	sequence := make([]T, end-start)
	for i := 0; i < end-start; i++ {
		sequence[i] = T(start + i)
	}
	return sequence
}

func RandomSequence[T Integer](start, end int) []T {
	sequence := Sequence[T](start, end)
	_ = shuffle(sequence)
	return sequence
}

func RandomSequenceFrom0[T Integer](end int) []T {
	return RandomSequence[T](0, end)
}

func randomN[T Integer](start, end int) func() T {
	sequence := Sequence[T](start, end)
	return func() T {
		return sequence[start+random.Intn(end-start)]
	}
}

func RandomSequenceRepeated[T Integer](start, end int) []T {
	sequence := make([]T, end-start)
	randN := randomN[T](start, end)
	for i := 0; i < end-start; i++ {
		sequence[i] = randN()
	}
	return sequence
}

func RandomSequenceRepeatedFrom0[T Integer](end int) []T {
	return RandomSequenceRepeated[T](0, end)
}
