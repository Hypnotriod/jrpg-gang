package domain

type UnitStats struct {
	Progress       UnitProgress       `json:"progress"`
	BaseAttributes UnitBaseAttributes `json:"baseAttributes"`
	Attributes     UnitAttributes     `json:"attributes"`
	Resistance     UnitResistance     `json:"resistance"`
}

func (s *UnitStats) TotalResistance() UnitResistance {
	resistance := s.Resistance
	resistance.IncreasePhysical(s.Attributes.Physique * RESISTANCE_MODIFICATION_FACTOR)
	return resistance
}
