package domain

import "jrpg-gang/util"

type UnitRecovery struct {
	UnitState
	Damage
}

func (s *UnitRecovery) Normalize() {
	s.Damage.Normalize()
	s.Health = util.Max(s.Health, 0)
	s.Stamina = util.Max(s.Stamina, 0)
	s.Mana = util.Max(s.Mana, 0)
	s.Stress = util.Max(s.Stress, 0)
}

func (s *UnitRecovery) MultiplyAll(factor float32) {
	s.Damage.MultiplyAll(factor)
	s.Health = util.MultiplyWithRounding(s.Health, factor)
	s.Stamina = util.MultiplyWithRounding(s.Stamina, factor)
	s.Mana = util.MultiplyWithRounding(s.Mana, factor)
	s.Stress = util.MultiplyWithRounding(s.Stress, factor)
}

func (s *UnitRecovery) Accumulate(state UnitRecovery) {
	s.Damage.Accumulate(state.Damage)
	s.UnitBaseAttributes.Accumulate(state.UnitBaseAttributes)
	s.Stress += state.Stress
}
