package util

type WeightedValue interface {
	Weight() int
}

func RandomPick[T WeightedValue](gen *RndGen, values []T) T {
	var sum int = 0
	for _, v := range values {
		sum += v.Weight()
	}
	limit := gen.rng.Int() % sum
	sum = 0
	for _, v := range values {
		sum += v.Weight()
		if limit < sum {
			return v
		}
	}
	return values[0]
}
