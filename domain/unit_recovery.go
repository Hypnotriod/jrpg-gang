package domain

import "jrpg-gang/util"

type UnitRecovery struct {
	UnitState
	Damage
}

func (r *UnitRecovery) Normalize() {
	r.Damage.Normalize()
	r.Health = util.Max(r.Health, 0)
	r.Stamina = util.Max(r.Stamina, 0)
	r.Mana = util.Max(r.Mana, 0)
	r.Stress = util.Max(r.Stress, 0)
}

func (r *UnitRecovery) MultiplyAll(factor float32) {
	r.Damage.MultiplyAll(factor)
	r.Health = util.MultiplyWithRounding(r.Health, factor)
	r.Stamina = util.MultiplyWithRounding(r.Stamina, factor)
	r.Mana = util.MultiplyWithRounding(r.Mana, factor)
	r.Stress = util.MultiplyWithRounding(r.Stress, factor)
}

func (r *UnitRecovery) EnchanceAll(value float32) {
	r.Damage.EnchanceAll(value)
	r.Health = util.AccumulateIfNotZeros(r.Health, value)
	r.Stamina = util.AccumulateIfNotZeros(r.Stamina, value)
	r.Mana = util.AccumulateIfNotZeros(r.Mana, value)
	r.Stress = util.AccumulateIfNotZeros(r.Stress, value)
}

func (r *UnitRecovery) Accumulate(state UnitRecovery) {
	r.Damage.Accumulate(state.Damage)
	r.UnitBaseAttributes.Accumulate(state.UnitBaseAttributes)
	r.Stress += state.Stress
}

func (r *UnitRecovery) HasEffect() bool {
	return r.Damage.HasEffect() ||
		r.Health != 0 ||
		r.Stamina != 0 ||
		r.Mana != 0 ||
		r.Stress != 0
}
