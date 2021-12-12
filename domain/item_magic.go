package domain

import "fmt"

type Magic struct {
	Item
	Requirements UnitAttributes           `json:"requirements"`
	UseCost      UnitBaseAttributes       `json:"useCost"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}

func (m Magic) String() string {
	return fmt.Sprintf(
		"%s, description: %s, requirements: {%v}, use cost: {%v}, damage: %v, modification: {%v}",
		m.Name,
		m.Description,
		m.Requirements,
		m.UseCost,
		m.Damage,
		m.Modification,
	)
}
