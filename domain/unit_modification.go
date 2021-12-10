package domain

import "fmt"

type UnitModification struct {
	BaseAttributes UnitBaseAttributes `json:"baseAttributes,omitempty"`
	Attributes     UnitAttributes     `json:"attributes,omitempty"`
	Resistance     UnitResistance     `json:"resistance,omitempty"`
	Damage         Damage             `json:"damage,omitempty"`
	Recovery       UnitState          `json:"recovery,omitempty"`
}

func (e *UnitModification) Accumulate(modification UnitModification) {
	e.BaseAttributes.Accumulate(modification.BaseAttributes)
	e.Attributes.Accumulate(modification.Attributes)
	e.Resistance.Accumulate(modification.Resistance)
	e.Damage.Accumulate(modification.Damage)
	e.Recovery.Accumulate(modification.Recovery)
}

func (e UnitModification) String() string {
	return fmt.Sprintf(
		"baseAttributes: {%v}, attributes: {%v}, resistance: {%v}, damage: {%v}, revovery: {%v}",
		e.BaseAttributes,
		e.Attributes,
		e.Resistance,
		e.Damage,
		e.Recovery,
	)
}
