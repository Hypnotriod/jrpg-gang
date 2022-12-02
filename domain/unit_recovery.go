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
	s.Fear = util.Max(s.Fear, 0)
	s.Curse = util.Max(s.Curse, 0)
}

func (s *UnitRecovery) Multiply(factor float32) {
	s.Damage.Multiply(factor)
	s.Health = util.MultiplyWithRounding(s.Health, factor)
	s.Stamina = util.MultiplyWithRounding(s.Stamina, factor)
	s.Mana = util.MultiplyWithRounding(s.Mana, factor)
	s.Fear = util.MultiplyWithRounding(s.Fear, factor)
	s.Curse = util.MultiplyWithRounding(s.Curse, factor)
}

func (s *UnitRecovery) Accumulate(state UnitRecovery) {
	s.Damage.Accumulate(state.Damage)
	s.UnitBaseAttributes.Accumulate(state.UnitBaseAttributes)
	s.Fear += state.Fear
	s.Curse += state.Curse
}
