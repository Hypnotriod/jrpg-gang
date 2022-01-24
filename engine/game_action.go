package engine

import (
	"fmt"
	"jrpg-gang/domain"
)

type GameAtionType string

const (
	GameAtionUse     GameAtionType = "use"
	GameAtionEquip   GameAtionType = "equip"
	GameAtionUnequip GameAtionType = "unequip"
	GameAtionMove    GameAtionType = "move"
)

type GameAction struct {
	Uid       uint            `json:"uid"`
	TargetUid uint            `json:"targetUid,omitempty"`
	ItemUid   uint            `json:"item_uid,omitempty"`
	Position  domain.Position `json:"position,omitempty"`
	Action    GameAtionType   `json:"action,omitempty"`
}

func (a GameAction) String() string {
	return fmt.Sprintf(
		"uid: %d, target uid: %d, item uid: %d, position: {%v}, action: {%v}",
		a.Uid,
		a.TargetUid,
		a.ItemUid,
		a.Position,
		a.Action,
	)
}
