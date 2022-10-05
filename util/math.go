package util

func AccumulateIfNotZeros[T Number](a, b T) T {
	if a != 0 && b != 0 {
		return a + b
	}
	return a
}

func MultiplyIfNotZeros[T Number](a, b T) T {
	if a != 0 && b != 0 {
		return a * b
	}
	return a
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Abs[T Number](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

func Round[T Float](v T) T {
	n := int64(v)
	if v > T(n)+0.5 {
		return T(n + 1)
	}
	return T(n)
}
