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
	Action    GameAtionType   `json:"action"`
	Uid       uint            `json:"uid"`
	TargetUid uint            `json:"targetUid,omitempty"`
	ItemUid   uint            `json:"item_uid,omitempty"`
	Position  domain.Position `json:"position,omitempty"`
}

func (a GameAction) String() string {
	return fmt.Sprintf(
		"%s, uid: %d, target uid: %d, item uid: %d, position: {%v}",
		a.Action,
		a.Uid,
		a.TargetUid,
		a.ItemUid,
		a.Position,
	)
}
