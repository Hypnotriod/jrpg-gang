package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Position) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.X, p.Y)
}

func (p *Position) Equals(position Position) bool {
	return p.X == position.X && p.Y == position.Y
}

func (p *Position) CheckActionRange(position Position, actionRange ActionRange) bool {
	minimum := util.AbsInt(position.X-p.X) >= actionRange.MinimumX && util.AbsInt(position.Y-p.Y) >= actionRange.MinimumY
	maximum := util.AbsInt(position.X-p.X) <= actionRange.MaximumX && util.AbsInt(position.Y-p.Y) <= actionRange.MaximumY
	return minimum && maximum
}
