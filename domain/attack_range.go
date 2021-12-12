package domain

import "fmt"

type AttackRange struct {
	Minimum float32 `json:"minimum,omitempty"`
	Maximum float32 `json:"maximum,omitempty"`
	Radius  float32 `json:"radius,omitempty"`
}

func (r AttackRange) String() string {
	return fmt.Sprintf("minimum: %g, maximum: %g, radius: %g",
		r.Minimum, r.Maximum, r.Radius)
}
