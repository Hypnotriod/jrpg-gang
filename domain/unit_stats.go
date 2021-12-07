package domain

import "fmt"

type UnitStats struct {
	Progress       UnitProgress       `json:"progress"`
	BaseAttributes UnitBaseAttributes `json:"baseAttributes"`
	Attributes     UnitAttributes     `json:"attributes"`
	Resistance     UnitResistance     `json:"resistance"`
}

func (s UnitStats) String() string {
	return fmt.Sprintf(
		"%v, attributes: {%v, %v}, resistance: {%v}",
		s.Progress,
		s.BaseAttributes,
		s.Attributes,
		s.Resistance,
	)
}
