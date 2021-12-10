package domain

import (
	"fmt"
)

type Unit struct {
	Name         string                   `json:"name"`
	State        UnitState                `json:"state"`
	Stats        UnitStats                `json:"stats"`
	Damage       []DamageImpact           `json:"damage"`
	Modification []UnitModificationImpact `json:"modification"`
	Inventory    UnitInventory            `json:"inventory"`
	Slots        map[EquipmentSlot]uint   `json:"slots"`
}

func (u Unit) String() string {
	return fmt.Sprintf(
		"%s, state: {%v}, stats: {%v}, damage: %v, modification: %v, inventory: {%v}, slots: %v",
		u.Name,
		u.State,
		u.Stats,
		u.Damage,
		u.Modification,
		u.Inventory,
		u.Slots,
	)
}

func (u *Unit) TotalAgility() float32 {
	var agility float32 = u.Stats.Attributes.Agility
	for _, ench := range u.Modification {
		agility += ench.Attributes.Agility
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			agility += ench.Attributes.Agility
		}
	}
	return agility
}

func (u *Unit) TotalIntelligence() float32 {
	var intelligence float32 = u.Stats.Attributes.Intelligence
	for _, ench := range u.Modification {
		intelligence += ench.Attributes.Intelligence
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			intelligence += ench.Attributes.Intelligence
		}
	}
	return intelligence
}

func (u *Unit) TotalLuck() float32 {
	var luck float32 = u.Stats.Attributes.Luck
	for _, ench := range u.Modification {
		luck += ench.Attributes.Luck
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			luck += ench.Attributes.Luck
		}
	}
	return luck
}

func (u *Unit) TotalInitiative() float32 {
	var initiative float32 = u.Stats.Attributes.Initiative
	for _, ench := range u.Modification {
		initiative += ench.Attributes.Initiative
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			initiative += ench.Attributes.Initiative
		}
	}
	return initiative
}

func (u *Unit) TotalModification() *UnitModification {
	var modification *UnitModification = &UnitModification{}
	for _, ench := range u.Modification {
		modification.Accumulate(ench.UnitModification)
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			modification.Accumulate(ench)
		}
	}
	return modification
}
