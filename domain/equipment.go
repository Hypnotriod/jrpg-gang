package domain

import "fmt"

type EquipmentSlot string

const (
	EquipmentSlotHead   EquipmentSlot = "head"
	EquipmentSlotNeck   EquipmentSlot = "neck"
	EquipmentSlotBody   EquipmentSlot = "body"
	EquipmentSlotHands  EquipmentSlot = "hands"
	EquipmentSlotLegs   EquipmentSlot = "legs"
	EquipmentSlotWeapon EquipmentSlot = "weapon"
)

type Equipment struct {
	Item
	Condition    float32        `json:"condition"`
	Slot         EquipmentSlot  `json:"slot"`
	Equipped     bool           `json:"equipped"`
	Requirements UnitAttributes `json:"requirements"`
}

func (e Equipment) String() string {
	return fmt.Sprintf(
		"Equipment: name: %s, type: %s, description: %s, slot: %s, equipped: %t, condition: %g, requirements: {%v}",
		e.Name, e.Type, e.Description, e.Slot, e.Equipped, e.Condition, e.Requirements)
}
