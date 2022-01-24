package engine

import (
	"fmt"
	"jrpg-gang/domain"
)

type GameUnit struct {
	domain.Unit
	FractionId uint   `json:"fractionId"`
	UserId     string `json:"userId,omitempty"`
}

func (u GameUnit) String() string {
	return fmt.Sprintf(
		"%v, fraction: %d, user id: %s",
		u.Unit,
		u.FractionId,
		u.UserId,
	)
}
