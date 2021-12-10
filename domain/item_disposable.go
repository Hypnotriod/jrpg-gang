package domain

import "fmt"

type Disposable struct {
	Item
	Quantity     uint                     `json:"quantity"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}

func (d Disposable) String() string {
	return fmt.Sprintf(
		"%s, description: %s, quantity: %d, modification: %v, damage: %v",
		d.Name,
		d.Description,
		d.Quantity,
		d.Modification,
		d.Damage,
	)
}
