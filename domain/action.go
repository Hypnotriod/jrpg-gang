package domain

type AtionType string

const (
	ActionUse        AtionType = "use"
	ActionEquip      AtionType = "equip"
	ActionUnequip    AtionType = "unequip"
	ActionPlace      AtionType = "place"
	ActionMove       AtionType = "move"
	ActionBuy        AtionType = "buy"
	ActionSell       AtionType = "sell"
	ActionActivate   AtionType = "activate"
	ActionDeactivate AtionType = "deactivate"
	ActionComplete   AtionType = "complete"
	ActionRepair     AtionType = "repair"
	ActionThrowAway  AtionType = "throwAway"
	ActionSkip       AtionType = "skip"
	ActionWait       AtionType = "wait"
	ActionLevelUp    AtionType = "levelUp"
	ActionSkillUp    AtionType = "skillUp"
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
	Uid       uint           `json:"uid,omitzero"`
	TargetUid uint           `json:"targetUid,omitzero"`
	ItemUid   uint           `json:"itemUid,omitzero"`
	Quantity  uint           `json:"quantity,omitzero"`
	Property  ActionProperty `json:"property,omitempty"`
	QuestCode QuestCode      `json:"questCode,omitempty"`
	Position  *Position      `json:"position,omitempty"`
}
