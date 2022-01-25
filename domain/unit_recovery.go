package domain

import "jrpg-gang/util"

type UnitRecovery struct {
	UnitState
}

func (s *UnitRecovery) Normalize() {
	s.Health = util.MaxFloat32(s.Health, 0)
	s.Stamina = util.MaxFloat32(s.Stamina, 0)
	s.Mana = util.MaxFloat32(s.Mana, 0)
	s.Fear = util.MaxFloat32(s.Fear, 0)
	s.Curse = util.MaxFloat32(s.Curse, 0)
}

func (s *UnitRecovery) Accumulate(state UnitRecovery) {
	s.UnitState.Accumulate(state.UnitState)
}
