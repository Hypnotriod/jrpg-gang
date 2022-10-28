package domain

import (
	"jrpg-gang/util"
)

type ActionRange struct {
	MinimumX int `json:"minimumX,omitempty"`
	MaximumX int `json:"maximumX,omitempty"`
	MinimumY int `json:"minimumY,omitempty"`
	MaximumY int `json:"maximumY,omitempty"`
	Radius   int `json:"radius,omitempty"`
}

func (r *ActionRange) Has() bool {
	return r.MaximumX != 0 ||
		r.MaximumY != 0 ||
		r.Radius != 0 ||
		r.MinimumX != 0 ||
		r.MinimumY != 0
}

func (r *ActionRange) Check(p1 Position, p2 Position) bool {
	minimum := util.Abs(p1.X-p2.X) >= r.MinimumX && util.Abs(p1.Y-p2.Y) >= r.MinimumY
	maximum := util.Abs(p1.X-p2.X) <= r.MaximumX && util.Abs(p1.Y-p2.Y) <= r.MaximumY
	return minimum && maximum
}
