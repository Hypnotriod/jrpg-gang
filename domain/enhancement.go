package domain

import "fmt"

type UnitEnhancement struct {
	BaseAttributes UnitBaseAttributes `json:"baseAttributes,omitempty"`
	Attributes     UnitAttributes     `json:"attributes,omitempty"`
	Resistance     UnitResistance     `json:"resistance,omitempty"`
	Damage         Damage             `json:"damage,omitempty"`
}

func (u *UnitEnhancement) Accumulate(enhancement UnitEnhancement) {
	u.BaseAttributes.Accumulate(enhancement.BaseAttributes)
	u.Attributes.Accumulate(enhancement.Attributes)
	u.Resistance.Accumulate(enhancement.Resistance)
	u.Damage.Accumulate(enhancement.Damage)
}

func (u *Unit) TotalEnhancement(checkEquipment bool) *UnitEnhancement {
	var enhancement *UnitEnhancement = &UnitEnhancement{}
	for _, e := range u.Enhancement {
		enhancement.Accumulate(e.UnitEnhancement)
	}
	if !checkEquipment {
		return enhancement
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

func (e UnitEnhancement) String() string {
	return fmt.Sprintf(
		"baseAttributes: {%v}, attributes: {%v}, resistance: {%v}, damage: {%v}",
		e.BaseAttributes,
		e.Attributes,
		e.Resistance,
		e.Damage,
	)
}
