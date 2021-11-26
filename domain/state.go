package domain

import "fmt"

type UnitState struct {
	UnitBaseAttributes
	Fear  float32 `json:"fear"`
	Curse float32 `json:"curse"`
}

func (s UnitState) String() string {
	return fmt.Sprintf(
		"%s, fear: %g, curse: %g",
		s.UnitBaseAttributes.String(),
		s.Fear,
		s.Curse,
	)
}
