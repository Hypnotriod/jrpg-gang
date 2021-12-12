package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type UnitBaseAttributes struct {
	Health  float32 `json:"health"`
	Stamina float32 `json:"stamina"`
	Mana    float32 `json:"mana"`
}

func (a UnitBaseAttributes) String() string {
	return fmt.Sprintf(
		"health: %g, stamina: %g, mana: %g",
		a.Health,
		a.Stamina,
		a.Mana,
	)
}

func (a *UnitBaseAttributes) Accumulate(attributes UnitBaseAttributes) {
	a.Health += attributes.Health
	a.Mana += attributes.Mana
	a.Stamina += attributes.Stamina
}

func (a *UnitBaseAttributes) Reduce(attributes UnitBaseAttributes) {
	a.Health -= attributes.Health
	a.Mana -= attributes.Mana
	a.Stamina -= attributes.Stamina
}

func (a *UnitBaseAttributes) Saturate(limit UnitBaseAttributes) {
	a.Health = util.MinFloat32(a.Health, limit.Health)
	a.Mana = util.MinFloat32(a.Mana, limit.Mana)
	a.Stamina = util.MinFloat32(a.Stamina, limit.Stamina)
}

func (a *UnitBaseAttributes) Normalize() {
	a.Health = util.MaxFloat32(a.Health, MINIMUM_BASE_ATTRIBUTE_HEALTH)
	a.Mana = util.MaxFloat32(a.Mana, MINIMUM_BASE_ATTRIBUTE_MANA)
	a.Stamina = util.MaxFloat32(a.Stamina, MINIMUM_BASE_ATTRIBUTE_STAMINA)
}
