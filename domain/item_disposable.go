package domain

import "fmt"

type Disposable struct {
	Item
	Damage      []DamageImpact          `json:"damage,omitempty"`
	Enhancement []UnitEnhancementImpact `json:"enhancement,omitempty"`
}

func (d Disposable) String() string {
	return fmt.Sprintf(
		"%s, description: %s, enhancement: %v, damage: %v",
		d.Name,
		d.Description,
		d.Enhancement,
		d.Damage,
	)
}
