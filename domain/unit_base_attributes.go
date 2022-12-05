package domain

import (
	"jrpg-gang/util"
)

type UnitBaseAttributes struct {
	Health  float32 `json:"health"`  // unit health
	Stamina float32 `json:"stamina"` // unit stamina
	Mana    float32 `json:"mana"`    // unit mana
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
	a.Health = util.Min(a.Health, limit.Health)
	a.Mana = util.Min(a.Mana, limit.Mana)
	a.Stamina = util.Min(a.Stamina, limit.Stamina)
}

func (a *UnitBaseAttributes) Normalize() {
	a.Health = util.Max(a.Health, MINIMUM_BASE_ATTRIBUTE_HEALTH)
	a.Mana = util.Max(a.Mana, MINIMUM_BASE_ATTRIBUTE_MANA)
	a.Stamina = util.Max(a.Stamina, MINIMUM_BASE_ATTRIBUTE_STAMINA)
}

func (a *UnitBaseAttributes) MultiplyAll(factor float32) {
	a.Health += util.MultiplyWithRounding(a.Health, factor)
	a.Mana += util.MultiplyWithRounding(a.Mana, factor)
	a.Stamina += util.MultiplyWithRounding(a.Stamina, factor)
}
