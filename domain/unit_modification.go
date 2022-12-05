package domain

type UnitModification struct {
	BaseAttributes UnitBaseAttributes `json:"baseAttributes,omitempty"`
	Attributes     UnitAttributes     `json:"attributes,omitempty"`
	Resistance     UnitResistance     `json:"resistance,omitempty"`
	Damage         Damage             `json:"damage,omitempty"`
	Recovery       UnitRecovery       `json:"recovery,omitempty"`
}

func (m *UnitModification) Accumulate(modification UnitModification) {
	m.BaseAttributes.Accumulate(modification.BaseAttributes)
	m.Attributes.Accumulate(modification.Attributes)
	m.Resistance.Accumulate(modification.Resistance)
	m.Damage.Accumulate(modification.Damage)
	m.Recovery.Accumulate(modification.Recovery)
}

func (m *UnitModification) MultiplyAll(factor float32) {
	m.BaseAttributes.MultiplyAll(factor)
	m.Attributes.MultiplyAll(factor)
	m.Resistance.MultiplyAll(factor)
	m.Damage.MultiplyAll(factor)
	m.Recovery.MultiplyAll(factor)
}
