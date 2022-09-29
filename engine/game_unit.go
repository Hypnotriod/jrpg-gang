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
	IsDead     bool            `json:"isDead,omitempty"`
	UserId     UserId          `json:"-"`
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
	return u.UserId != UserIdEmpty
}

func (u *GameUnit) Clone() *GameUnit {
	r := &GameUnit{}
	r.Faction = u.Faction
	r.IsDead = u.IsDead
	r.UserId = u.UserId
	r.Unit = *u.Unit.Clone()
	if u.PlayerInfo != nil {
		r.PlayerInfo = u.PlayerInfo.Clone()
	}
	return r
}
