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

type UnitStats struct {
	Progress       UnitProgress       `json:"progress"`
	BaseAttributes UnitBaseAttributes `json:"baseAttributes"`
	Attributes     UnitAttributes     `json:"attributes"`
	Resistance     UnitResistance     `json:"resistance"`
}

func (s UnitStats) String() string {
	return fmt.Sprintf(
		"%v, attributes: {%v, %v}, resistance: {%v}",
		s.Progress,
		s.BaseAttributes,
		s.Attributes,
		s.Resistance,
	)
}

type UnitState struct {
	UnitBaseAttributes
	Fear  float32 `json:"fear"`
	Curse float32 `json:"curse"`
}

type UnitProgress struct {
	Level      float32 `json:"level"`
	Experience float32 `json:"experience"`
}

func (p UnitProgress) String() string {
	return fmt.Sprintf("level: %g, experience: %g",
		p.Level, p.Experience)
}

func (u *Unit) TotalAgility(checkEquipment bool) float32 {
	var agility float32 = u.Stats.Attributes.Agility
	for _, e := range u.Enhancement {
		agility += e.Attributes.Agility
	}
	if !checkEquipment {
		return agility
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

func (u *Unit) TotalLuck(checkEquipment bool) float32 {
	var luck float32 = u.Stats.Attributes.Luck
	for _, e := range u.Enhancement {
		luck += e.Attributes.Luck
	}
	if !checkEquipment {
		return luck
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
