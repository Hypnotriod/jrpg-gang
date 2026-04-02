package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameUnitFaction uint

const (
	GameUnitFactionLeft  GameUnitFaction = 0
	GameUnitFactionRight GameUnitFaction = 1
)

type GameUnit struct {
	domain.Unit
	Faction       GameUnitFaction       `json:"faction"`
	PlayerInfo    *PlayerInfo           `json:"playerInfo,omitempty"`
	Drop          []domain.UnitBooty    `json:"drop,omitempty"`
	QuestTriggers []domain.QuestTrigger `json:"questTriggers,omitempty"`
	IsDead        bool                  `json:"isDead,omitempty"`
}

func NewGameUnit(unit *domain.Unit) *GameUnit {
	u := &GameUnit{}
	u.Unit = *unit
	return u
}

func (u *GameUnit) HasPlayerId() bool {
	return u.PlayerInfo != nil && u.PlayerInfo.Id != PlayerIdEmpty
}

func (u *GameUnit) GetPlayerId() PlayerId {
	if u.PlayerInfo != nil {
		return u.PlayerInfo.Id
	}
	return PlayerIdEmpty
}

func (u *GameUnit) Clone() *GameUnit {
	r := &GameUnit{}
	r.Faction = u.Faction
	r.IsDead = u.IsDead
	r.Drop = []domain.UnitBooty{}
	r.Drop = append(r.Drop, u.Drop...)
	r.QuestTriggers = util.Map(u.QuestTriggers, func(trigger domain.QuestTrigger) domain.QuestTrigger {
		return *trigger.Clone()
	})
	r.Unit = *u.Unit.Clone()
	if u.PlayerInfo != nil {
		r.PlayerInfo = u.PlayerInfo.Clone()
	}
	return r
}

func (u *GameUnit) ToPersist() *GameUnit {
	r := u.Clone()
	r.Inventory.FillDescriptor()
	return r
}

func (u *GameUnit) PrepareForUser() {
	u.Stats.Progress.ExperienceNext = u.NextLevelExp()
}
