package domain

import (
	"jrpg-gang/util"
)

type UnitState struct {
	UnitBaseAttributes
	Stress       float32 `json:"stress"`                 // reduces action chance, affects retreat chance
	IsStunned    bool    `json:"isStunned,omitempty"`    // stun flag
	WaitingOrder uint    `json:"waitingOrder,omitempty"` // waiting order
}

func (s *UnitState) RestoreDefault(limit UnitBaseAttributes) {
	s.Stress = 0
	s.ActionPoints = 0
	s.Health = limit.Health
	s.Mana = limit.Mana
	s.Stamina = limit.Stamina
}

func (s *UnitState) RestoreToHalf(limit UnitBaseAttributes) {
	s.Stress = max(s.Stress-HALF_CHANCE, s.Stress)
	s.Health = max(util.Round(limit.Health/2), s.Health)
	s.Mana = max(util.Round(limit.Mana/2), s.Mana)
	s.Stamina = max(util.Round(limit.Stamina/2), s.Stamina)
}

func (s *UnitState) Accumulate(state UnitState) {
	s.UnitBaseAttributes.Accumulate(state.UnitBaseAttributes)
	s.Stress -= state.Stress
}

func (s *UnitState) Normalize() {
	s.UnitBaseAttributes.Normalize()
	s.Stress = max(s.Stress, 0)
}

func (s *UnitState) Saturate(limit UnitBaseAttributes) {
	s.UnitBaseAttributes.Saturate(limit)
	s.Stress = min(MAXIMUM_CHANCE, s.Stress)
}
