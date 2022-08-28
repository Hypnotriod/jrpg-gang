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

func Find[T any](values []T, predicate func(value *T) bool) *T {
	for i := range values {
		if predicate(&values[i]) {
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

func Any[T any](values []T, predicate func(value *T) bool) bool {
	for i := range values {
		if predicate(&values[i]) {
			return true
		}
	}
	return false
}

func Anyp[T any](values []*T, predicate func(value *T) bool) bool {
	for i := range values {
		if predicate(values[i]) {
			return true
		}
	}
	return false
}

func Every[T any](values []T, predicate func(value *T) bool) bool {
	if len(values) == 0 {
		return false
	}
	for i := range values {
		if !predicate(&values[i]) {
			return false
		}
	}
	return true
}

func Everyp[T any](values []*T, predicate func(value *T) bool) bool {
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

func Filter[T any](values []T, predicate func(value *T) bool) []T {
	result := []T{}
	for i := range values {
		if predicate(&values[i]) {
			result = append(result, values[i])
		}
	}
	return result
}

func Filterp[T any](values []*T, predicate func(value *T) bool) []*T {
	result := []*T{}
	for i := range values {
		if predicate(values[i]) {
			result = append(result, values[i])
		}
	}
	return result
}

func Map[T any, J any](values []T, mapper func(value *T) J) []J {
	result := make([]J, 0, len(values))
	for i := range values {
		result = append(result, mapper(&values[i]))
	}
	return result
}

func Mapp[T any, J any](values []*T, mapper func(value *T) *J) []*J {
	result := []*J{}
	for i := range values {
		result = append(result, mapper(values[i]))
	}
	return result
}
