package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type UnitSlots map[EquipmentSlot]uint

type Unit struct {
	Uid          uint                     `json:"uid"`
	Name         string                   `json:"name"`
	State        UnitState                `json:"state"`
	Stats        UnitStats                `json:"stats"`
	Damage       []DamageImpact           `json:"damage"`
	Modification []UnitModificationImpact `json:"modification"`
	Inventory    UnitInventory            `json:"inventory"`
	Slots        UnitSlots                `json:"slots"`
	Position     Position                 `json:"position"`
}

func (u Unit) String() string {
	return fmt.Sprintf(
		"%s, state: {%v}, stats: {%v}, damage: [%s], modification: [%s], inventory: {%v}, slots: %v, position: {%v}",
		u.Name,
		u.State,
		u.Stats,
		util.AsCommaSeparatedSlice(u.Damage),
		util.AsCommaSeparatedSlice(u.Modification),
		u.Inventory,
		u.Slots,
		u.Position,
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
	return util.MaxFloat32(agility, 0)
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
	return util.MaxFloat32(intelligence, 0)
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
	return util.MaxFloat32(luck, 0)
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
	return util.MaxFloat32(initiative, 0)
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

func (u *Unit) CheckRequirements(requirements UnitAttributes) bool {
	attributes := u.TotalModification().Attributes
	attributes.Accumulate(u.Stats.Attributes)
	attributes.Normalize()
	return attributes.Strength >= requirements.Strength &&
		attributes.Physique >= requirements.Physique &&
		attributes.Agility >= requirements.Agility &&
		attributes.Endurance >= requirements.Endurance &&
		attributes.Intelligence >= requirements.Intelligence &&
		attributes.Initiative >= requirements.Initiative &&
		attributes.Luck >= requirements.Luck
}

func (u *Unit) CheckUseCost(useCost UnitBaseAttributes) bool {
	return u.State.Health >= useCost.Health &&
		u.State.Mana >= useCost.Mana &&
		u.State.Stamina >= useCost.Stamina
}
