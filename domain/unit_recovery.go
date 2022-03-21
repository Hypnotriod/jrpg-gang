package domain

import "jrpg-gang/util"

type UnitRecovery struct {
	UnitState
}

func (s *UnitRecovery) Normalize() {
	s.Health = util.Max(s.Health, 0)
	s.Stamina = util.Max(s.Stamina, 0)
	s.Mana = util.Max(s.Mana, 0)
	s.Fear = util.Max(s.Fear, 0)
	s.Curse = util.Max(s.Curse, 0)
}

func (s *UnitRecovery) Accumulate(state UnitRecovery) {
	s.UnitState.Accumulate(state.UnitState)
}
