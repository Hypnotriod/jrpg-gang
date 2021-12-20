package domain

import "fmt"

type Ammunition struct {
	Item
	Kind     string         `json:"kind"`
	Quantity uint           `json:"quantity"`
	Damage   []DamageImpact `json:"damage,omitempty"`
}

func (a Ammunition) String() string {
	return fmt.Sprintf(
		"%s, description: %s, kind: %s, quantity: %d, damage: %v",
		a.Name,
		a.Description,
		a.Kind,
		a.Quantity,
		a.Damage,
	)
}
