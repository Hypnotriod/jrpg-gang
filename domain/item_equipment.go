package domain

import (
	"fmt"
	"jrpg-gang/util"
)

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
	Wearout      float32            `json:"wearout"`
	Durability   float32            `json:"durability"`
	Slot         EquipmentSlot      `json:"slot"`
	SlotsNumber  uint               `json:"slotsNumber"`
	Equipped     bool               `json:"equipped,omitempty"`
	Requirements UnitAttributes     `json:"requirements"`
	Modification []UnitModification `json:"modification"`
}

func (e Equipment) String() string {
	return fmt.Sprintf(
		"%s, type: %s, description: %s, slot: %s, slots: %d, equipped: %t, wearout: %g, durability: %g, requirements: {%v}, modification: [%s], uid: %d",
		e.Name,
		e.Type,
		e.Description,
		e.Slot,
		e.SlotsNumber,
		e.Equipped,
		e.Wearout,
		e.Durability,
		e.Requirements,
		util.AsCommaSeparatedObjectsSlice(e.Modification),
		e.Uid,
	)
}

func (e *Equipment) IncreaseWearout() {
	e.Wearout++
}

func (e *Equipment) IsBroken() bool {
	return e.Wearout >= e.Durability
}
