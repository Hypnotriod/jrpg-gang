package domain

import "fmt"

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
