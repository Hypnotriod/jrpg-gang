package domain

import "jrpg-gang/util"

type UnitStats struct {
	Progress       UnitProgress       `json:"progress" bson:"progress"`
	BaseAttributes UnitBaseAttributes `json:"baseAttributes" bson:"baseAttributes"`
	Attributes     UnitAttributes     `json:"attributes" bson:"attributes"`
	Resistance     UnitResistance     `json:"resistance" bson:"resistance"`
}

func (s *UnitStats) TotalResistance() UnitResistance {
	resistance := s.Resistance
	physiqueBased := util.Floor(s.Attributes.Physique * RESISTANCE_MODIFICATION_FACTOR)
	resistance.AccumulatePhysical(physiqueBased)
	return resistance
}
