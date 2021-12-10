package domain

import "fmt"

type Disposable struct {
	Item
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}

func (d Disposable) String() string {
	return fmt.Sprintf(
		"%s, description: %s, modification: %v, damage: %v",
		d.Name,
		d.Description,
		d.Modification,
		d.Damage,
	)
}
