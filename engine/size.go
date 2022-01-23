package engine

import "fmt"

type Size struct {
	Width  uint
	Height uint
}

func (s Size) String() string {
	return fmt.Sprintf("width: %d, height: %d", s.Width, s.Height)
}

func (s *Size) Equals(size *Size) bool {
	return s.Width == size.Width && s.Height == size.Height
}
