package domain

import "fmt"

type Weapon struct {
	Equipment
	Hands       int               `json:"hands"`
	Damage      []DamageImpact    `json:"damage"`
	Enhancement []UnitEnhancement `json:"enhancement"`
}

func (w Weapon) String() string {
	return fmt.Sprintf(
		"Weapon: name: %s, description: %s, condition: %g, equipped: %t, hands: %d, damage: %s, requirements: {%v}, enhancement: %v",
		w.Name,
		w.Description,
		w.Condition,
		w.Equipped,
		w.Hands,
		w.Damage,
		w.Requirements,
		w.Enhancement)
}
