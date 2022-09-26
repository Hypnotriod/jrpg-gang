package engine

import (
	"fmt"
	"jrpg-gang/domain"
)

type Spot struct {
	Name        string             `json:"name"`
	Code        string             `json:"code"`
	Battlefield Battlefield        `json:"battlefield"`
	Booty       []domain.UnitBooty `json:"booty"`
}

func (s Spot) String() string {
	return fmt.Sprintf(
		"%s, code: %s, battlefield: {%v}",
		s.Name,
		s.Code,
		s.Battlefield,
	)
}
