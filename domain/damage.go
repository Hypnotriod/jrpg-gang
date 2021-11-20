package domain

import (
	"fmt"
	"strings"
)

type Damage struct {
	Stabbing float32 `json:"stabbing,omitempty"`
	Cutting  float32 `json:"cutting,omitempty"`
	Crushing float32 `json:"crushing,omitempty"`
	Fire     float32 `json:"fire,omitempty"`
	Cold     float32 `json:"cold,omitempty"`
	Lighting float32 `json:"lighting,omitempty"`
	Fear     float32 `json:"fear,omitempty"`
	Poison   float32 `json:"poison,omitempty"`
	Curse    float32 `json:"curse,omitempty"`
	Stunning float32 `json:"stunning,omitempty"`
	// todo
}

func (d Damage) String() string {
	props := []string{}

	if d.Stabbing != 0 {
		props = append(props, fmt.Sprintf("stabbing: %g", d.Stabbing))
	}
	if d.Cutting != 0 {
		props = append(props, fmt.Sprintf("cutting: %g", d.Cutting))
	}
	if d.Crushing != 0 {
		props = append(props, fmt.Sprintf("crushing: %g", d.Crushing))
	}
	if d.Fire != 0 {
		props = append(props, fmt.Sprintf("fire: %g", d.Fire))
	}
	if d.Cold != 0 {
		props = append(props, fmt.Sprintf("cold: %g", d.Cold))
	}
	if d.Lighting != 0 {
		props = append(props, fmt.Sprintf("lighting: %g", d.Lighting))
	}
	if d.Fear != 0 {
		props = append(props, fmt.Sprintf("fear: %g", d.Fear))
	}
	if d.Poison != 0 {
		props = append(props, fmt.Sprintf("poison: %g", d.Poison))
	}
	if d.Curse != 0 {
		props = append(props, fmt.Sprintf("curse: %g", d.Curse))
	}
	if d.Stunning != 0 {
		props = append(props, fmt.Sprintf("stunning: %g", d.Stunning))
	}

	return strings.Join(props, ", ")
}
