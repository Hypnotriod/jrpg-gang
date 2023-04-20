package domain

import (
	"jrpg-gang/util"
)

type UnitBaseAttributes struct {
	Health       float32 `json:"health" bson:"health"`             // unit health
	Stamina      float32 `json:"stamina" bson:"stamina"`           // unit stamina
	Mana         float32 `json:"mana" bson:"mana"`                 // unit mana
	ActionPoints float32 `json:"actionPoints" bson:"actionPoints"` // action points
}

func (a *UnitBaseAttributes) Accumulate(attributes UnitBaseAttributes) {
	a.Health += attributes.Health
	a.Mana += attributes.Mana
	a.Stamina += attributes.Stamina
	a.ActionPoints += attributes.ActionPoints
}

func (a *UnitBaseAttributes) Reduce(attributes UnitBaseAttributes) {
	a.Health -= attributes.Health
	a.Mana -= attributes.Mana
	a.Stamina -= attributes.Stamina
	a.ActionPoints -= attributes.ActionPoints
}

func (a *UnitBaseAttributes) Saturate(limit UnitBaseAttributes) {
	a.Health = util.Min(a.Health, limit.Health)
	a.Mana = util.Min(a.Mana, limit.Mana)
	a.Stamina = util.Min(a.Stamina, limit.Stamina)
	a.ActionPoints = util.Min(a.ActionPoints, limit.ActionPoints)
}

func (a *UnitBaseAttributes) Normalize() {
	a.Health = util.Max(a.Health, MINIMUM_BASE_ATTRIBUTE_HEALTH)
	a.Mana = util.Max(a.Mana, MINIMUM_BASE_ATTRIBUTE_MANA)
	a.Stamina = util.Max(a.Stamina, MINIMUM_BASE_ATTRIBUTE_STAMINA)
	a.ActionPoints = util.Max(a.ActionPoints, MINIMUM_BASE_ATTRIBUTE_ACTION_POINTS)
}

func (a *UnitBaseAttributes) MultiplyAll(factor float32) {
	a.Health = util.MultiplyWithRounding(a.Health, factor)
	a.Mana = util.MultiplyWithRounding(a.Mana, factor)
	a.Stamina = util.MultiplyWithRounding(a.Stamina, factor)
	a.ActionPoints = util.MultiplyWithRounding(a.ActionPoints, factor)
}

func (a *UnitBaseAttributes) EnchanceAll(value float32) {
	a.Health = util.AccumulateIfNotZeros(a.Health, value)
	a.Mana = util.AccumulateIfNotZeros(a.Mana, value)
	a.Stamina = util.AccumulateIfNotZeros(a.Stamina, value)
	a.ActionPoints = util.AccumulateIfNotZeros(a.ActionPoints, value)
}

func (a *UnitBaseAttributes) ReduceActionPoint() {
	if a.ActionPoints > 0 {
		a.ActionPoints--
	}
}

func (a *UnitBaseAttributes) ClearActionPoints() {
	a.ActionPoints = 0
}
