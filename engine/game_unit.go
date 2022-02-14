package engine

import (
	"fmt"
	"jrpg-gang/domain"
)

type UserId string
type GameUnitClass string

const (
	UnitClassTank  GameUnitClass = "tank"
	UnitClassRogue GameUnitClass = "rogue"
	UnitClassMage  GameUnitClass = "mage"
)

type GameUnit struct {
	domain.Unit
	FractionId uint   `json:"fractionId"`
	UserId     UserId `json:"userId,omitempty"`
}

func (u GameUnit) String() string {
	return fmt.Sprintf(
		"%v, fraction: %d, user id: %s",
		u.Unit,
		u.FractionId,
		u.UserId,
	)
}
