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

func (s *UnitRecovery) Accumulate(state UnitRecovery) {
	s.Damage.Accumulate(state.Damage)
	s.UnitBaseAttributes.Accumulate(state.UnitBaseAttributes)
	s.Fear += state.Fear
	s.Curse += state.Curse
}
