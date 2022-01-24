package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type ActionRange struct {
	MinimumX int `json:"minimumX,omitempty"`
	MaximumX int `json:"maximumX,omitempty"`
	MinimumY int `json:"minimumY,omitempty"`
	MaximumY int `json:"maximumY,omitempty"`
	RadiusX  int `json:"radiusX,omitempty"`
	RadiusY  int `json:"radiusY,omitempty"`
}

func (r ActionRange) String() string {
	return fmt.Sprintf("minimum: (%d:%d), maximum: (%d:%d), radius: (%d:%d)",
		r.MinimumX,
		r.MinimumY,
		r.MaximumX,
		r.MaximumY,
		r.RadiusX,
		r.RadiusY,
	)
}

func (r *ActionRange) Check(p1 Position, p2 Position) bool {
	minimum := util.AbsInt(p1.X-p2.X) >= r.MinimumX && util.AbsInt(p1.Y-p2.Y) >= r.MinimumY
	maximum := util.AbsInt(p1.X-p2.X) <= r.MaximumX && util.AbsInt(p1.Y-p2.Y) <= r.MaximumY
	return minimum && maximum
}
