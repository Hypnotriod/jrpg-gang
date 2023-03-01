package domain

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
	Requirements UnitRequirements   `json:"requirements"`
	Modification []UnitModification `json:"modification"`
}

func (e *Equipment) IncreaseWearout() {
	if e.Durability != 0 {
		e.Wearout++
	}
}

func (e *Equipment) IsBroken() bool {
	return e.Durability != 0 && e.Wearout >= e.Durability
}
