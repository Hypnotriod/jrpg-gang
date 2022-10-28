package domain

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Position) Equals(position Position) bool {
	return p.X == position.X && p.Y == position.Y
}
