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

func Find[T any](values []T, test func(value *T) bool) *T {
	for i := range values {
		if test(&values[i]) {
			return &values[i]
		}
	}
	return nil
}

func Findp[T any](values []*T, test func(value *T) bool) *T {
	for i := range values {
		if test(values[i]) {
			return values[i]
		}
	}
	return nil
}

func Filter[T any](values []T, test func(value *T) bool) []*T {
	result := []*T{}
	for i := range values {
		if test(&values[i]) {
			result = append(result, &values[i])
		}
	}
	return result
}

func Filterp[T any](values []*T, test func(value *T) bool) []*T {
	result := []*T{}
	for i := range values {
		if test(values[i]) {
			result = append(result, values[i])
		}
	}
	return result
}
