package domain

import (
	"fmt"
)

type UnitBaseAttributes struct {
	Health  float32 `json:"health"`
	Stamina float32 `json:"stamina"`
	Mana    float32 `json:"mana"`
}

func (a *UnitBaseAttributes) Accumulate(attributes UnitBaseAttributes) {
	a.Health += attributes.Health
	a.Mana += attributes.Mana
	a.Stamina += attributes.Stamina
}

func (a UnitBaseAttributes) String() string {
	return fmt.Sprintf(
		"health: %g, stamina: %g, mana: %g",
		a.Health,
		a.Stamina,
		a.Mana,
	)
}
