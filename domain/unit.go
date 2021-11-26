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
	Items       []interface{}           `json:"items"`
}

func (u Unit) String() string {
	return fmt.Sprintf(
		"Unit: name: %s, state: {%v}, stats: {%v}, impact: %v, enhancement: %v, items: %v",
		u.Name,
		u.State,
		u.Stats,
		u.Impact,
		u.Enhancement,
		u.Items,
	)
}

func (u *Unit) TotalAgility() float32 {
	var agility float32 = u.Stats.Attributes.Agility
	for _, e := range u.Enhancement {
		agility += e.Attributes.Agility
	}
	for _, item := range u.Items {
		equipment, ok := AsEquipment(item)
		if !ok || !equipment.Equipped {
			continue
		}
		for _, e := range equipment.Enhancement {
			agility += e.Attributes.Agility
		}
	}
	return agility
}

func (u *Unit) TotalLuck() float32 {
	var luck float32 = u.Stats.Attributes.Luck
	for _, e := range u.Enhancement {
		luck += e.Attributes.Luck
	}
	for _, item := range u.Items {
		equipment, ok := AsEquipment(item)
		if !ok || !equipment.Equipped {
			continue
		}
		for _, e := range equipment.Enhancement {
			luck += e.Attributes.Luck
		}
	}
	return luck
}

func (u *Unit) TotalEnhancement() *UnitEnhancement {
	var enhancement *UnitEnhancement = &UnitEnhancement{}
	for _, e := range u.Enhancement {
		enhancement.Accumulate(e.UnitEnhancement)
	}
	for _, item := range u.Items {
		equipment, ok := AsEquipment(item)
		if !ok || !equipment.Equipped {
			continue
		}
		for _, e := range equipment.Enhancement {
			enhancement.Accumulate(e)
		}
	}
	return enhancement
}
