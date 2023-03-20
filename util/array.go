package util

func Contains[T Ordered](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func IndexOf[T Ordered](values []T, value T) int {
	for i, v := range values {
		if v == value {
			return i
		}
	}
	return -1
}

func Clone[T any](values []T) []T {
	accumulator := make([]T, 0, len(values))
	return append(accumulator, values...)
}

func Shuffle[T any](gen *RndGen, values []T) []T {
	for i1 := range values {
		i2 := gen.rng.Int() % len(values)
		values[i1], values[i2] = values[i2], values[i1]
	}
	return values
}

func Find[T any](values []T, predicate func(value T) bool) *T {
	for i := range values {
		if predicate(values[i]) {
			return &values[i]
		}
	}
	return nil
}

func Findp[T any](values []*T, predicate func(value *T) bool) *T {
	for i := range values {
		if predicate(values[i]) {
			return values[i]
		}
	}
	return nil
}

func Any[T any](values []T, predicate func(value T) bool) bool {
	for i := range values {
		if predicate(values[i]) {
			return true
		}
	}
	return false
}

func Every[T any](values []T, predicate func(value T) bool) bool {
	if len(values) == 0 {
		return false
	}
	for i := range values {
		if !predicate(values[i]) {
			return false
		}
	}
	return true
}

func Filter[T any](values []T, predicate func(value T) bool) []T {
	result := []T{}
	for i := range values {
		if predicate(values[i]) {
			result = append(result, values[i])
		}
	}
	return result
}

func Reduce[T, J any](values []T, accumulator J, reducer func(accumulator J, value T) J) J {
	for i := range values {
		accumulator = reducer(accumulator, values[i])
	}
	return accumulator
}

func Map[T any, J any](values []T, mapper func(value T) J) []J {
	result := make([]J, 0, len(values))
	for i := range values {
		result = append(result, mapper(values[i]))
	}
	return result
}
