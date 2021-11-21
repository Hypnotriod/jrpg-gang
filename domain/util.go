package domain

func maxFloat32(a, b float32) float32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func minFloat32(a, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}
