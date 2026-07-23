package util

func OrElse[T any](ptr *T, def *T) *T {
	if ptr == nil {
		return def
	}
	return ptr
}

func OrNew[T any](ptr *T) *T {
	if ptr == nil {
		return new(T)
	}
	return ptr
}
