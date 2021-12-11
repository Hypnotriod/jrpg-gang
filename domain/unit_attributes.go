package domain

import (
	"fmt"
	"strings"
)

type UnitAttributes struct {
	Strength     float32 `json:"strength"`
	Physique     float32 `json:"physique"`
	Agility      float32 `json:"agility"`
	Endurance    float32 `json:"endurance"`
	Intelligence float32 `json:"intelligence"`
	Initiative   float32 `json:"initiative"`
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
	if a.Agility != 0 {
		props = append(props, fmt.Sprintf("agility: %g", a.Agility))
	}
	if a.Endurance != 0 {
		props = append(props, fmt.Sprintf("endurance: %g", a.Endurance))
	}
	if a.Intelligence != 0 {
		props = append(props, fmt.Sprintf("intelligence: %g", a.Intelligence))
	}
	if a.Initiative != 0 {
		props = append(props, fmt.Sprintf("initiative: %g", a.Initiative))
	}
	if a.Luck != 0 {
		props = append(props, fmt.Sprintf("luck: %g", a.Luck))
	}
	return strings.Join(props, ", ")
}

func (a *UnitAttributes) Accumulate(attributes UnitAttributes) {
	a.Strength += attributes.Strength
	a.Physique += attributes.Physique
	a.Agility += attributes.Agility
	a.Endurance += attributes.Endurance
	a.Intelligence += attributes.Intelligence
	a.Initiative += attributes.Initiative
	a.Luck += attributes.Luck
}
