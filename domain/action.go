package domain

type AtionType string

const (
	ActionUse     AtionType = "use"
	ActionEquip   AtionType = "equip"
	ActionUnequip AtionType = "unequip"
	ActionPlace   AtionType = "place"
	ActionMove    AtionType = "move"
	ActionBuy     AtionType = "buy"
	ActionSkip    AtionType = "skip"
)

type Action struct {
	Action    AtionType `json:"action"`
	Uid       uint      `json:"uid,omitempty"`
	TargetUid uint      `json:"targetUid,omitempty"`
	ItemUid   uint      `json:"itemUid,omitempty"`
	Quantity  uint      `json:"quantity,omitempty"`
	Position  *Position `json:"position,omitempty"`
}
