package domain

type UnitModification struct {
	BaseAttributes UnitBaseAttributes `json:"baseAttributes,omitzero"`
	Attributes     UnitAttributes     `json:"attributes,omitzero"`
	Resistance     UnitResistance     `json:"resistance,omitzero"`
	Damage         Damage             `json:"damage,omitzero"`
	Recovery       UnitRecovery       `json:"recovery,omitzero"`
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

func (m *UnitModification) EnchanceAll(value float32) {
	m.BaseAttributes.EnchanceAll(value)
	m.Attributes.EnchanceAll(value)
	m.Resistance.EnchanceAll(value)
	m.Damage.EnchanceAll(value)
	m.Recovery.EnchanceAll(value)
}
