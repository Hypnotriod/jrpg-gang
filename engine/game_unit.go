package engine

import (
	"fmt"
	"jrpg-gang/domain"
)

type UserId string
type GameUnitClass string

const (
	UserIdEmpty UserId = ""
)

const (
	UnitClassTank  GameUnitClass = "tank"
	UnitClassRogue GameUnitClass = "rogue"
	UnitClassMage  GameUnitClass = "mage"
)

type GameUnitFaction uint

const (
	GameUnitFactionLeft  GameUnitFaction = 0
	GameUnitFactionRight GameUnitFaction = 1
)

type GameUnit struct {
	domain.Unit
	Faction GameUnitFaction `json:"faction"`
	UserId  UserId          `json:"userId,omitempty"`
}

func (u GameUnit) String() string {
	return fmt.Sprintf(
		"%v, faction: %d, user id: %s",
		u.Unit,
		u.Faction,
		u.UserId,
	)
}

func (u *GameUnit) HasUserId() bool {
	return len(u.UserId) != 0
}
