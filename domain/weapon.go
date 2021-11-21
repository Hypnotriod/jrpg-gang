package domain

import "fmt"

type Weapon struct {
	Equipment
	Damage []DamageImpact `json:"damage"`
}

func (w Weapon) String() string {
	return fmt.Sprintf(
		"Weapon: name: %s, description: %s, condition: %g, equipped: %t, slots: %d, damage: %s, requirements: {%v}, enhancement: %v",
		w.Name,
		w.Description,
		w.Condition,
		w.Equipped,
		w.SlotsNumber,
		w.Damage,
		w.Requirements,
		w.Enhancement)
}
