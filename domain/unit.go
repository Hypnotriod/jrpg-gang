package domain

import (
	"fmt"
)

type Unit struct {
	Name        string                  `json:"name"`
	State       UnitState               `json:"state"`
	Stats       UnitStats               `json:"stats"`
	Impact      []DamageImpact          `json:"impact"`
	Enhancement []UnitEnhancementImpact `json:"enhancement"`
	Inventory   Inventory               `json:"inventory"`
}

func (u Unit) String() string {
	return fmt.Sprintf(
		"%s, state: {%v}, stats: {%v}, impact: %v, enhancement: %v, inventory: {%v}",
		u.Name,
		u.State,
		u.Stats,
		u.Impact,
		u.Enhancement,
		u.Inventory,
	)
}

func (u *Unit) Equipment(equippedOnly bool) []*Equipment {
	equipment := []*Equipment{}
	for i := range u.Inventory.Armor {
		item := &u.Inventory.Armor[i].Equipment
		if item.Equipped || !equippedOnly {
			equipment = append(equipment, item)
		}
	}
	for i := range u.Inventory.Weapon {
		item := &u.Inventory.Weapon[i].Equipment
		if item.Equipped || !equippedOnly {
			equipment = append(equipment, item)
		}
	}
	return equipment
}

func (u *Unit) TotalAgility() float32 {
	var agility float32 = u.Stats.Attributes.Agility
	for _, ench := range u.Enhancement {
		agility += ench.Attributes.Agility
	}
	for _, item := range u.Equipment(true) {
		for _, ench := range item.Enhancement {
			agility += ench.Attributes.Agility
		}
	}
	return agility
}

func (u *Unit) TotalIntelligence() float32 {
	var intelligence float32 = u.Stats.Attributes.Intelligence
	for _, ench := range u.Enhancement {
		intelligence += ench.Attributes.Intelligence
	}
	for _, item := range u.Equipment(true) {
		for _, ench := range item.Enhancement {
			intelligence += ench.Attributes.Intelligence
		}
	}
	return intelligence
}

func (u *Unit) TotalLuck() float32 {
	var luck float32 = u.Stats.Attributes.Luck
	for _, ench := range u.Enhancement {
		luck += ench.Attributes.Luck
	}
	for _, item := range u.Equipment(true) {
		for _, ench := range item.Enhancement {
			luck += ench.Attributes.Luck
		}
	}
	return luck
}

func (u *Unit) TotalInitiative() float32 {
	var initiative float32 = u.Stats.Attributes.Initiative
	for _, ench := range u.Enhancement {
		initiative += ench.Attributes.Initiative
	}
	for _, item := range u.Equipment(true) {
		for _, ench := range item.Enhancement {
			initiative += ench.Attributes.Initiative
		}
	}
	return initiative
}

func (u *Unit) TotalEnhancement() *UnitEnhancement {
	var enhancement *UnitEnhancement = &UnitEnhancement{}
	for _, ench := range u.Enhancement {
		enhancement.Accumulate(ench.UnitEnhancement)
	}
	for _, item := range u.Equipment(true) {
		for _, ench := range item.Enhancement {
			enhancement.Accumulate(ench)
		}
	}
	return enhancement
}
