package domain

import "fmt"

type Weapon struct {
	Equipment
	Impact []DamageImpact `json:"impact"`
}

func (w Weapon) String() string {
	return fmt.Sprintf(
		"%s, description: %s, condition: %g, equipped: %t, slots: %d, impact: %v, requirements: {%v}, enhancement: {%v}",
		w.Name,
		w.Description,
		w.Condition,
		w.Equipped,
		w.SlotsNumber,
		w.Impact,
		w.Requirements,
		w.Enhancement,
	)
}
