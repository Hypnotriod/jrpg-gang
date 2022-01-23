package engine

import "fmt"

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Position) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.X, p.Y)
}

func (p *Position) Equals(position *Position) bool {
	return p.X == position.X && p.Y == position.Y
}
