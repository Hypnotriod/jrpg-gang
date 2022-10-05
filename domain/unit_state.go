package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type UnitState struct {
	UnitBaseAttributes
	Fear      float32 `json:"fear"`                // affects retreat chance
	Curse     float32 `json:"curse"`               // reduces action chance
	IsStunned bool    `json:"isStunned,omitempty"` // stun flag
}

func (s *UnitState) IsEmpty() bool {
	return s.Fear == 0 && s.Curse == 0 && s.Health == 0 && s.Mana == 0 && s.Stamina == 0
}

func (s *UnitState) RestoreToHalf(limit UnitBaseAttributes) {
	s.Curse = util.Max(s.Curse-HALF_CHANCE, s.Curse)
	s.Fear = util.Max(s.Fear-HALF_CHANCE, s.Fear)
	s.Health = util.Max(util.Round(limit.Health/2), s.Health)
	s.Mana = util.Max(util.Round(limit.Mana/2), s.Mana)
	s.Stamina = util.Max(util.Round(limit.Stamina/2), s.Stamina)
}

func (s *UnitState) Accumulate(state UnitState) {
	s.UnitBaseAttributes.Accumulate(state.UnitBaseAttributes)
	s.Fear -= state.Fear
	s.Curse -= state.Curse
}

func (s *UnitState) Normalize() {
	s.UnitBaseAttributes.Normalize()
	s.Fear = util.Max(s.Fear, 0)
	s.Curse = util.Max(s.Curse, 0)
}

func (s *UnitState) Saturate(limit UnitBaseAttributes) {
	s.UnitBaseAttributes.Saturate(limit)
	s.Curse = util.Min(MAXIMUM_CHANCE, s.Curse)
	s.Fear = util.Min(MAXIMUM_CHANCE, s.Fear)
}

func (s UnitState) String() string {
	return fmt.Sprintf(
		"%s, fear: %g, curse: %g",
		s.UnitBaseAttributes.String(),
		s.Fear,
		s.Curse,
	)
}
