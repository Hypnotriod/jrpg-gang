package engine

type Size struct {
	Width  uint
	Height uint
}

func (s *Size) Equals(size *Size) bool {
	return s.Width == size.Width && s.Height == size.Height
}
