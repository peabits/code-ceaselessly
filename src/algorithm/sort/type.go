package sort

import (
	"cmp"
)

// cmp.Ordered
type Ordered interface {
	cmp.Ordered
	Number | ~string
}

type Number interface {
	Integer | Float
}

type Integer interface {
	Signed | Unsigned
}

// byte = ~int32
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// rune = ~uint8
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}
