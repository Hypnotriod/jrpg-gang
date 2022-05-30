package util

type Integer interface {
	int | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

type Ordered interface {
	Number | ~string
}
