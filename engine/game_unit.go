package engine

import (
	"fmt"
	"jrpg-gang/domain"
)

type GameUnit struct {
	domain.Unit
	Position   Position `json:"position"`
	FractionId uint     `json:"fractionId"`
	UserId     string   `json:"userId,omitempty"`
}

func (u GameUnit) String() string {
	return fmt.Sprintf(
		"%v, position: {%v}, fraction: %d, user id: %s",
		u.Unit,
		u.Position,
		u.FractionId,
		u.UserId,
	)
}
