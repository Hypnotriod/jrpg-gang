package util

func Contains[T Ordered](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
