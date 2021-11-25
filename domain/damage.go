package domain

import (
	"fmt"
	"strings"

	"jrpg-gang/util"
)

type Damage struct {
	Stabbing   float32 `json:"stabbing,omitempty"`
	Cutting    float32 `json:"cutting,omitempty"`
	Crushing   float32 `json:"crushing,omitempty"`
	Fire       float32 `json:"fire,omitempty"`
	Cold       float32 `json:"cold,omitempty"`
	Lighting   float32 `json:"lighting,omitempty"`
	Poison     float32 `json:"poison,omitempty"`
	Stunning   float32 `json:"stunning,omitempty"`
	Exhaustion float32 `json:"exhaustion,omitempty"`
	Bleeding   float32 `json:"bleeding,omitempty"`
	Fear       float32 `json:"fear,omitempty"`
	Curse      float32 `json:"curse,omitempty"`
	IsCritical bool
}

func (d *Damage) Accumulate(damage Damage) {
	d.Stabbing += damage.Stabbing
	d.Cutting += damage.Cutting
	d.Crushing += damage.Crushing
	d.Fire += damage.Fire
	d.Cold += damage.Cold
	d.Lighting += damage.Lighting
	d.Poison += damage.Poison
	d.Stunning += damage.Stunning
	d.Exhaustion += damage.Exhaustion
	d.Bleeding += damage.Bleeding
	d.Fear += damage.Fear
	d.Curse += damage.Curse
}

func (d *Damage) Reduce(damage Damage) {
	d.Stabbing -= damage.Stabbing
	d.Cutting -= damage.Cutting
	d.Crushing -= damage.Crushing
	d.Fire -= damage.Fire
	d.Cold -= damage.Cold
	d.Lighting -= damage.Lighting
	d.Poison -= damage.Poison
	d.Stunning -= damage.Stunning
	d.Exhaustion -= damage.Exhaustion
	d.Bleeding -= damage.Bleeding
	d.Fear -= damage.Fear
	d.Curse -= damage.Curse
}

func (d *Damage) HasEffect() bool {
	return d.Stabbing != 0 ||
		d.Cutting != 0 ||
		d.Crushing != 0 ||
		d.Fire != 0 ||
		d.Cold != 0 ||
		d.Lighting != 0 ||
		d.Poison != 0 ||
		d.Stunning != 0 ||
		d.Exhaustion != 0 ||
		d.Bleeding != 0 ||
		d.Fear != 0 ||
		d.Curse != 0
}

func (d *Damage) Enchance(attributes UnitAttributes, damage Damage) {
	d.Stabbing = util.AccumulateIfNotZerosFloat32(d.Stabbing, attributes.Strength)
	d.Cutting = util.AccumulateIfNotZerosFloat32(d.Cutting, attributes.Strength)
	d.Crushing = util.AccumulateIfNotZerosFloat32(d.Crushing, attributes.Strength)
	d.Fire = util.AccumulateIfNotZerosFloat32(d.Fire, attributes.Intelligence)
	d.Cold = util.AccumulateIfNotZerosFloat32(d.Cold, attributes.Intelligence)
	d.Lighting = util.AccumulateIfNotZerosFloat32(d.Lighting, attributes.Intelligence)
	d.Stunning = util.AccumulateIfNotZerosFloat32(d.Stunning, attributes.Strength)
	d.Exhaustion = util.AccumulateIfNotZerosFloat32(d.Exhaustion, attributes.Intelligence)
	d.Bleeding = util.AccumulateIfNotZerosFloat32(d.Bleeding, attributes.Strength)

	d.Stabbing = util.AccumulateIfNotZerosFloat32(d.Stabbing, damage.Stabbing)
	d.Cutting = util.AccumulateIfNotZerosFloat32(d.Cutting, damage.Cutting)
	d.Crushing = util.AccumulateIfNotZerosFloat32(d.Crushing, damage.Crushing)
	d.Fire = util.AccumulateIfNotZerosFloat32(d.Fire, damage.Fire)
	d.Cold = util.AccumulateIfNotZerosFloat32(d.Cold, damage.Cold)
	d.Lighting = util.AccumulateIfNotZerosFloat32(d.Lighting, damage.Lighting)
	d.Poison = util.AccumulateIfNotZerosFloat32(d.Poison, damage.Poison)
	d.Stunning = util.AccumulateIfNotZerosFloat32(d.Stunning, damage.Stunning)
	d.Exhaustion = util.AccumulateIfNotZerosFloat32(d.Exhaustion, damage.Exhaustion)
	d.Bleeding = util.AccumulateIfNotZerosFloat32(d.Bleeding, damage.Bleeding)
	d.Fear = util.AccumulateIfNotZerosFloat32(d.Fear, damage.Fear)
	d.Curse = util.AccumulateIfNotZerosFloat32(d.Curse, damage.Curse)
}

func (d *Damage) Multiply(factor float32) {
	d.Stabbing = util.MultiplyIfNotZerosFloat32(d.Stabbing, factor)
	d.Cutting = util.MultiplyIfNotZerosFloat32(d.Cutting, factor)
	d.Crushing = util.MultiplyIfNotZerosFloat32(d.Crushing, factor)
	d.Fire = util.MultiplyIfNotZerosFloat32(d.Fire, factor)
	d.Cold = util.MultiplyIfNotZerosFloat32(d.Cold, factor)
	d.Lighting = util.MultiplyIfNotZerosFloat32(d.Lighting, factor)
	d.Poison = util.MultiplyIfNotZerosFloat32(d.Poison, factor)
	d.Stunning = util.MultiplyIfNotZerosFloat32(d.Stunning, factor)
	d.Exhaustion = util.MultiplyIfNotZerosFloat32(d.Exhaustion, factor)
	d.Bleeding = util.MultiplyIfNotZerosFloat32(d.Bleeding, factor)
	d.Fear = util.MultiplyIfNotZerosFloat32(d.Fear, factor)
	d.Curse = util.MultiplyIfNotZerosFloat32(d.Curse, factor)
}

func (d *Damage) Normalize() {
	d.Stabbing = util.MaxFloat32(d.Stabbing, 0)
	d.Cutting = util.MaxFloat32(d.Cutting, 0)
	d.Crushing = util.MaxFloat32(d.Crushing, 0)
	d.Fire = util.MaxFloat32(d.Fire, 0)
	d.Cold = util.MaxFloat32(d.Cold, 0)
	d.Lighting = util.MaxFloat32(d.Lighting, 0)
	d.Poison = util.MaxFloat32(d.Poison, 0)
	d.Stunning = util.MaxFloat32(d.Stunning, 0)
	d.Exhaustion = util.MaxFloat32(d.Exhaustion, 0)
	d.Bleeding = util.MaxFloat32(d.Bleeding, 0)
	d.Fear = util.MaxFloat32(d.Fear, 0)
	d.Curse = util.MaxFloat32(d.Curse, 0)
}

func (d *Damage) Apply(state *UnitState) {
	state.Health -= d.Stabbing + d.Cutting + d.Crushing + d.Fire + d.Cold + d.Lighting + d.Poison + d.Bleeding
	state.Stamina -= d.Stunning
	state.Mana -= d.Exhaustion
	state.Curse += d.Curse
	state.Fear += d.Fear

	state.Health = util.MaxFloat32(state.Health, 0)
	state.Stamina = util.MaxFloat32(state.Stamina, 0)
	state.Mana = util.MaxFloat32(state.Mana, 0)
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
	if d.Poison != 0 {
		props = append(props, fmt.Sprintf("poison: %g", d.Poison))
	}
	if d.Stunning != 0 {
		props = append(props, fmt.Sprintf("stunning: %g", d.Stunning))
	}
	if d.Exhaustion != 0 {
		props = append(props, fmt.Sprintf("exhaustion: %g", d.Exhaustion))
	}
	if d.Bleeding != 0 {
		props = append(props, fmt.Sprintf("bleeding: %g", d.Bleeding))
	}
	if d.Fear != 0 {
		props = append(props, fmt.Sprintf("fear: %g", d.Fear))
	}
	if d.Curse != 0 {
		props = append(props, fmt.Sprintf("curse: %g", d.Curse))
	}
	if d.IsCritical {
		props = append(props, "critical!")
	}

	return strings.Join(props, ", ")
}
