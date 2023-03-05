package domain

type AtionType string

const (
	ActionUse       AtionType = "use"
	ActionEquip     AtionType = "equip"
	ActionUnequip   AtionType = "unequip"
	ActionPlace     AtionType = "place"
	ActionMove      AtionType = "move"
	ActionBuy       AtionType = "buy"
	ActionSell      AtionType = "sell"
	ActionRepair    AtionType = "repair"
	ActionThrowAway AtionType = "throwAway"
	ActionSkip      AtionType = "skip"
	ActionLevelUp   AtionType = "levelUp"
	ActionSkillUp   AtionType = "skillUp"
)

type ActionProperty string

const (
	AtionPropertyStrength     ActionProperty = "strength"
	AtionPropertyPhysique     ActionProperty = "physique"
	AtionPropertyAgility      ActionProperty = "agility"
	AtionPropertyEndurance    ActionProperty = "endurance"
	AtionPropertyIntelligence ActionProperty = "intelligence"
	AtionPropertyInitiative   ActionProperty = "initiative"
	AtionPropertyLuck         ActionProperty = "luck"
	AtionPropertyHealth       ActionProperty = "health"
	AtionPropertyStamina      ActionProperty = "stamina"
	AtionPropertyMana         ActionProperty = "mana"
)

type Action struct {
	Action    AtionType      `json:"action"`
	Uid       uint           `json:"uid,omitempty"`
	TargetUid uint           `json:"targetUid,omitempty"`
	ItemUid   uint           `json:"itemUid,omitempty"`
	Quantity  uint           `json:"quantity,omitempty"`
	Property  ActionProperty `json:"property,omitempty"`
	Position  *Position      `json:"position,omitempty"`
}
