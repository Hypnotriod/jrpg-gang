package util

func AccumulateIfNotZerosFloat32(a, b float32) float32 {
	if a != 0 && b != 0 {
		return a + b
	} else {
		return a
	}
}

func MultiplyIfNotZerosFloat32(a, b float32) float32 {
	if a != 0 && b != 0 {
		return a * b
	} else {
		return a
	}
}

func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func AbsInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
