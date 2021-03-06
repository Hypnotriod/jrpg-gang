package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type Magic struct {
	Item
	Requirements UnitAttributes           `json:"requirements"`
	Range        ActionRange              `json:"range"`
	UseCost      UnitBaseAttributes       `json:"useCost"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}

func (m Magic) String() string {
	return fmt.Sprintf(
		"%s, description: %s, requirements: {%v}, use cost: {%v}, range: {%v}, damage: [%s], modification: [%s], uid: %d",
		m.Name,
		m.Description,
		m.Requirements,
		m.UseCost,
		m.Range,
		util.AsCommaSeparatedObjectsSlice(m.Damage),
		util.AsCommaSeparatedObjectsSlice(m.Modification),
		m.Uid,
	)
}
