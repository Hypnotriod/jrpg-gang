package domain

import (
	"fmt"
	"strings"
)

type Unit struct {
	Name        string                  `json:"name"`
	State       UnitState               `json:"state"`
	Stats       UnitStats               `json:"stats"`
	Impact      []DamageImpact          `json:"impact"`
	Enhancement []UnitEnhancementImpact `json:"enhancement"`
	Items       []interface{}           `json:"items"`
}

func (u Unit) String() string {
	return fmt.Sprintf(
		"Unit: name: %s, state: {%v}, stats: {%v}, impact: %v, enhancement: %v, items: %v",
		u.Name,
		u.State,
		u.Stats,
		u.Impact,
		u.Enhancement,
		u.Items,
	)
}

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

type UnitBaseAttributes struct {
	Health  float32 `json:"health"`
	Stamina float32 `json:"stamina"`
	Mana    float32 `json:"mana"`
}

func (a UnitBaseAttributes) String() string {
	props := []string{}
	if a.Health != 0 {
		props = append(props, fmt.Sprintf("health: %g", a.Health))
	}
	if a.Stamina != 0 {
		props = append(props, fmt.Sprintf("stamina: %g", a.Stamina))
	}
	if a.Mana != 0 {
		props = append(props, fmt.Sprintf("mana: %g", a.Mana))
	}
	return strings.Join(props, ", ")
}

type UnitState struct {
	UnitBaseAttributes
	Fear  float32 `json:"fear"`
	Curse float32 `json:"curse"`
}

type UnitProgress struct {
	Level      float32 `json:"level"`
	Experience float32 `json:"experience"`
}

func (p UnitProgress) String() string {
	return fmt.Sprintf("level: %g, experience: %g",
		p.Level, p.Experience)
}

type UnitAttributes struct {
	Strength     float32 `json:"strength"`
	Physique     float32 `json:"physique"`
	Dexterity    float32 `json:"dexterity"`
	Endurance    float32 `json:"endurance"`
	Intelligence float32 `json:"intelligence"`
	Luck         float32 `json:"luck"`
}

func (a UnitAttributes) String() string {
	props := []string{}
	if a.Strength != 0 {
		props = append(props, fmt.Sprintf("strength: %g", a.Strength))
	}
	if a.Physique != 0 {
		props = append(props, fmt.Sprintf("physique: %g", a.Physique))
	}
	if a.Dexterity != 0 {
		props = append(props, fmt.Sprintf("dexterity: %g", a.Dexterity))
	}
	if a.Endurance != 0 {
		props = append(props, fmt.Sprintf("endurance: %g", a.Endurance))
	}
	if a.Intelligence != 0 {
		props = append(props, fmt.Sprintf("intelligence: %g", a.Intelligence))
	}
	if a.Luck != 0 {
		props = append(props, fmt.Sprintf("luck: %g", a.Luck))
	}
	return strings.Join(props, ", ")
}
