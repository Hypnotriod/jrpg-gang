package domain

import "fmt"

type EquipmentSlot string

const (
	EquipmentSlotHead   EquipmentSlot = "head"
	EquipmentSlotNeck   EquipmentSlot = "neck"
	EquipmentSlotBody   EquipmentSlot = "body"
	EquipmentSlotHands  EquipmentSlot = "hand"
	EquipmentSlotLegs   EquipmentSlot = "leg"
	EquipmentSlotWeapon EquipmentSlot = "weapon"
)

type Equipment struct {
	Item
	Condition    float32           `json:"condition"`
	Slot         EquipmentSlot     `json:"slot"`
	SlotsNumber  uint              `json:"slotsNumber"`
	Equipped     bool              `json:"equipped"`
	Requirements UnitAttributes    `json:"requirements"`
	Enhancement  []UnitEnhancement `json:"enhancement"`
}

func (e Equipment) String() string {
	return fmt.Sprintf(
		"Equipment: name: %s, type: %s, description: %s, slot: %s, slots: %d, equipped: %t, condition: %g, requirements: {%v}, enhancement: %v",
		e.Name, e.Type, e.Description, e.Slot, e.SlotsNumber, e.Equipped, e.Condition, e.Requirements, e.Enhancement)
}
