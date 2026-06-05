package domain

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var PositionOutOfBounds = Position{X: -1, Y: -1}

func (p *Position) Equals(position Position) bool {
	return p.X == position.X && p.Y == position.Y
}

func (p *Position) Shift(position Position) {
	p.X += position.X
	p.Y += position.Y
}
