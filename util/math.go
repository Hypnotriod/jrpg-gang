package util

func AccumulateIfNotZeros[T Number](a, b T) T {
	if a != 0 && b != 0 {
		return a + b
	}
	return a
}

func MultiplyWithRounding[T Float](a, b T) T {
	return T(uint64((a * b) + 0.5))
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
	if n < 0 {
		if v < T(n)-0.5 {
			return T(n - 1)
		}
		return T(n)
	}
	if v > T(n)+0.5 {
		return T(n + 1)
	}
	return T(n)
}

func Floor[T Float](v T) T {
	n := int64(v)
	if T(n) == v {
		return v
	}
	if n < 0 {
		return T(n - 1)
	}
	return T(n)
}

func Ceil[T Float](v T) T {
	n := int64(v)
	if T(n) == v {
		return v
	}
	if n < 0 {
		return T(n)
	}
	return T(n + 1)
}
