package domain

import "jrpg-gang/util"

type UnitStats struct {
	Progress       UnitProgress       `json:"progress"`
	BaseAttributes UnitBaseAttributes `json:"baseAttributes"`
	Attributes     UnitAttributes     `json:"attributes"`
	Resistance     UnitResistance     `json:"resistance"`
}

func (s *UnitStats) TotalResistance() UnitResistance {
	resistance := s.Resistance
	physiqueBased := util.Floor(s.Attributes.Physique * RESISTANCE_MODIFICATION_FACTOR)
	resistance.AccumulatePhysical(physiqueBased)
	return resistance
}
