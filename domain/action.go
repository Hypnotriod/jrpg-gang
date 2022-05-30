package domain

import (
	"fmt"
)

type AtionType string

const (
	ActionUse     AtionType = "use"
	ActionEquip   AtionType = "equip"
	ActionUnequip AtionType = "unequip"
	ActionPlace   AtionType = "place"
	ActionMove    AtionType = "move"
	ActionBuy     AtionType = "buy"
)

type Action struct {
	Action    AtionType `json:"action"`
	Uid       uint      `json:"uid,omitempty"`
	TargetUid uint      `json:"targetUid,omitempty"`
	ItemUid   uint      `json:"itemUid,omitempty"`
	Position  *Position `json:"position,omitempty"`
}

func (a Action) String() string {
	return fmt.Sprintf(
		"%s, uid: %d, target uid: %d, item uid: %d, position: {%v}",
		a.Action,
		a.Uid,
		a.TargetUid,
		a.ItemUid,
		a.Position,
	)
}
