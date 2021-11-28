package domain

import "fmt"

type UnitEnhancement struct {
	BaseAttributes UnitBaseAttributes `json:"baseAttributes,omitempty"`
	Attributes     UnitAttributes     `json:"attributes,omitempty"`
	Resistance     UnitResistance     `json:"resistance,omitempty"`
	Damage         Damage             `json:"damage,omitempty"`
	Recovery       UnitState          `json:"recovery,omitempty"`
}

func (e *UnitEnhancement) Accumulate(enhancement UnitEnhancement) {
	e.BaseAttributes.Accumulate(enhancement.BaseAttributes)
	e.Attributes.Accumulate(enhancement.Attributes)
	e.Resistance.Accumulate(enhancement.Resistance)
	e.Damage.Accumulate(enhancement.Damage)
	e.Recovery.Accumulate(enhancement.Recovery)
}

func (e UnitEnhancement) String() string {
	return fmt.Sprintf(
		"baseAttributes: {%v}, attributes: {%v}, resistance: {%v}, damage: {%v}, revovery: {%v}",
		e.BaseAttributes,
		e.Attributes,
		e.Resistance,
		e.Damage,
		e.Recovery,
	)
}
