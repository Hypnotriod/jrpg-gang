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
	Poison   float32 `json:"poison,omitempty"`
	Fear     float32 `json:"fear,omitempty"`
	Curse    float32 `json:"curse,omitempty"`
	Stunning float32 `json:"stunning,omitempty"`
}

func (r *Damage) Accumulate(damage *Damage) {
	r.Stabbing += damage.Stabbing
	r.Cutting += damage.Cutting
	r.Crushing += damage.Crushing
	r.Fire += damage.Fire
	r.Cold += damage.Cold
	r.Lighting += damage.Lighting
	r.Poison += damage.Poison
	r.Fear += damage.Fear
	r.Curse += damage.Curse
	r.Stunning += damage.Stunning
}

func (r *Damage) Reduce(damage *Damage) {
	r.Stabbing -= damage.Stabbing
	r.Cutting -= damage.Cutting
	r.Crushing -= damage.Crushing
	r.Fire -= damage.Fire
	r.Cold -= damage.Cold
	r.Lighting -= damage.Lighting
	r.Poison -= damage.Poison
	r.Fear -= damage.Fear
	r.Curse -= damage.Curse
	r.Stunning -= damage.Stunning
}

func (r *Damage) Normalize() {
	r.Stabbing = maxFloat32(r.Stabbing, 0)
	r.Cutting = maxFloat32(r.Cutting, 0)
	r.Crushing = maxFloat32(r.Crushing, 0)
	r.Fire = maxFloat32(r.Fire, 0)
	r.Cold = maxFloat32(r.Cold, 0)
	r.Lighting = maxFloat32(r.Lighting, 0)
	r.Poison = maxFloat32(r.Poison, 0)
	r.Fear = maxFloat32(r.Fear, 0)
	r.Curse = maxFloat32(r.Curse, 0)
	r.Stunning = maxFloat32(r.Stunning, 0)
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
