package domain

import "fmt"

type UnitModification struct {
	BaseAttributes UnitBaseAttributes `json:"baseAttributes,omitempty"`
	Attributes     UnitAttributes     `json:"attributes,omitempty"`
	Resistance     UnitResistance     `json:"resistance,omitempty"`
	Damage         Damage             `json:"damage,omitempty"`
	Recovery       UnitState          `json:"recovery,omitempty"`
}

func (m UnitModification) String() string {
	return fmt.Sprintf(
		"baseAttributes: {%v}, attributes: {%v}, resistance: {%v}, damage: {%v}, revovery: {%v}",
		m.BaseAttributes,
		m.Attributes,
		m.Resistance,
		m.Damage,
		m.Recovery,
	)
}

func (m *UnitModification) Accumulate(modification UnitModification) {
	m.BaseAttributes.Accumulate(modification.BaseAttributes)
	m.Attributes.Accumulate(modification.Attributes)
	m.Resistance.Accumulate(modification.Resistance)
	m.Damage.Accumulate(modification.Damage)
	m.Recovery.Accumulate(modification.Recovery)
}
