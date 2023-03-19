package util

func CloneMap[K Ordered, V any](src map[K]V) map[K]V {
	dts := make(map[K]V, len(src))
	for key, value := range src {
		dts[key] = value
	}
	return dts
}
