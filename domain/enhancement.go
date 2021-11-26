package domain

import "fmt"

type UnitEnhancement struct {
	BaseAttributes UnitBaseAttributes `json:"baseAttributes,omitempty"`
	Attributes     UnitAttributes     `json:"attributes,omitempty"`
	Resistance     UnitResistance     `json:"resistance,omitempty"`
	Damage         Damage             `json:"damage,omitempty"`
}

func (e *UnitEnhancement) Accumulate(enhancement UnitEnhancement) {
	e.BaseAttributes.Accumulate(enhancement.BaseAttributes)
	e.Attributes.Accumulate(enhancement.Attributes)
	e.Resistance.Accumulate(enhancement.Resistance)
	e.Damage.Accumulate(enhancement.Damage)
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
