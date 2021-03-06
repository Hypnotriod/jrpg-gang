package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type UnitState struct {
	UnitBaseAttributes
	Fear  float32 `json:"fear"`
	Curse float32 `json:"curse"`
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

func (s UnitState) String() string {
	return fmt.Sprintf(
		"%s, fear: %g, curse: %g",
		s.UnitBaseAttributes.String(),
		s.Fear,
		s.Curse,
	)
}
