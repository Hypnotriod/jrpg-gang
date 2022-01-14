package domain

import (
	"fmt"
	"jrpg-gang/util"
)

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
		"%s, description: %s, requirements: {%v}, use cost: {%v}, range: {%v}, damage: [%s], modification: [%s]",
		m.Name,
		m.Description,
		m.Requirements,
		m.UseCost,
		m.Range,
		util.AsCommaSeparatedSlice(m.Damage),
		util.AsCommaSeparatedSlice(m.Modification),
	)
}
