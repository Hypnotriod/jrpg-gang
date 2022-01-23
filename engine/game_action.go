package engine

type GameAtionType string

const (
	GameAtionUseItem GameAtionType = "useItem"
	GameAtionEquip   GameAtionType = "equip"
	GameAtionUnequip GameAtionType = "unequip"
)

type GameAction struct {
	Uid       uint          `json:"uid,omitempty"`
	TargetUid uint          `json:"targetUid,omitempty"`
	Position  Position      `json:"position,omitempty"`
	Action    GameAtionType `json:"action,omitempty"`
}
