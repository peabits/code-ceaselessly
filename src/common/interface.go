package common

import "cmp"

type measurable interface {
	Len() int
}

func Len(v measurable) int {
	return v.Len()
}

type maximum[T cmp.Ordered] interface {
	Max() T
}

type minimum[T cmp.Ordered] interface {
	Min() T
}

func Max[T cmp.Ordered](v maximum[T]) T {
	return v.Max()
}

func Min[T cmp.Ordered](v minimum[T]) T {
	return v.Min()
}

type limit[T cmp.Ordered] interface {
	maximum[T]
	minimum[T]
}

func MinMax[T cmp.Ordered](v limit[T]) (T, T) {
	return v.Min(), v.Max()
}

