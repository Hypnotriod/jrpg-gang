package domain

import "fmt"

type Weapon struct {
	Equipment
	Damage []DamageImpact `json:"damage"`
}

func (w Weapon) String() string {
	return fmt.Sprintf(
		"%s, description: %s, wearout: %g, durability: %g, equipped: %t, slots: %d, damage: %v, requirements: {%v}, enhancement: {%v}",
		w.Name,
		w.Description,
		w.Wearout,
		w.Durability,
		w.Equipped,
		w.SlotsNumber,
		w.Damage,
		w.Requirements,
		w.Enhancement,
	)
}
