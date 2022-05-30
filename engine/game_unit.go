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
	Faction    GameUnitFaction `json:"faction"`
	PlayerInfo *PlayerInfo     `json:"playerInfo,omitempty"`
	userId     UserId
}

func (u GameUnit) String() string {
	return fmt.Sprintf(
		"%v, faction: %d, user id: %s",
		u.Unit,
		u.Faction,
		u.userId,
	)
}

func (u *GameUnit) HasUserId() bool {
	return u.userId != UserIdEmpty
}

func (u *GameUnit) GetUserId() UserId {
	return u.userId
}

func (u *GameUnit) SetUserId(userId UserId) {
	u.userId = userId
}
