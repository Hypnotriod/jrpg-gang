package domain

import "fmt"

type Weapon struct {
	Equipment
	Range   AttackRange        `json:"range"`
	UseCost UnitBaseAttributes `json:"useCost"`
	Damage  []DamageImpact     `json:"damage"`
}

func (w Weapon) String() string {
	return fmt.Sprintf(
		"%s, description: %s, wearout: %g, durability: %g, equipped: %t, slots: %d, requirements: {%v}, use cost: {%v}, range: {%v}, damage: %v, modification: {%v}",
		w.Name,
		w.Description,
		w.Wearout,
		w.Durability,
		w.Equipped,
		w.SlotsNumber,
		w.Requirements,
		w.UseCost,
		w.Range,
		w.Damage,
		w.Modification,
	)
}
