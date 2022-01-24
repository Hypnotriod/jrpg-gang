package engine

type GameAtionType string

const (
	GameAtionUse     GameAtionType = "use"
	GameAtionEquip   GameAtionType = "equip"
	GameAtionUnequip GameAtionType = "unequip"
	GameAtionMove    GameAtionType = "move"
)

type GameAction struct {
	Uid       uint          `json:"uid"`
	TargetUid uint          `json:"targetUid,omitempty"`
	ItemUid   uint          `json:"item_uid,omitempty"`
	Position  Position      `json:"position,omitempty"`
	Action    GameAtionType `json:"action,omitempty"`
}
