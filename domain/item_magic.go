package domain

import "fmt"

type Magic struct {
	Item
	Requirements UnitAttributes           `json:"requirements"`
	Range        AttackRange              `json:"range"`
	UseCost      UnitBaseAttributes       `json:"useCost"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}

func (m Magic) String() string {
	return fmt.Sprintf(
		"%s, description: %s, requirements: {%v}, use cost: {%v}, range: {%v}, damage: %v, modification: {%v}",
		m.Name,
		m.Description,
		m.Requirements,
		m.UseCost,
		m.Range,
		m.Damage,
		m.Modification,
	)
}
