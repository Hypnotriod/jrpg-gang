package engine

import "jrpg-gang/domain"

type PlayerId string

const (
	PlayerIdEmpty PlayerId = ""
)

type PlayerInfo struct {
	Id        PlayerId         `json:"playerId"`
	Nickname  string           `json:"nickname"`
	Class     domain.UnitClass `json:"class"`
	Code      domain.UnitCode  `json:"code"`
	Level     uint             `json:"level"`
	UnitUid   uint             `json:"unitUid,omitzero"`
	IsOffline bool             `json:"isOffline,omitzero"`
	IsReady   bool             `json:"isReady,omitzero"`
	IsGuest   bool             `json:"isGuest,omitzero"`
}

func (p PlayerInfo) Clone() *PlayerInfo {
	return &p
}
