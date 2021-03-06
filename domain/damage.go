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
	Exhaustion float32 `json:"exhaustion,omitempty"`
	ManaDrain  float32 `json:"manaDrain,omitempty"`
	Bleeding   float32 `json:"bleeding,omitempty"`
	Fear       float32 `json:"fear,omitempty"`
	Curse      float32 `json:"curse,omitempty"`
	IsCritical bool    `json:"isCritical,omitempty"`
}

func (d *Damage) Accumulate(damage Damage) {
	d.Stabbing += damage.Stabbing
	d.Cutting += damage.Cutting
	d.Crushing += damage.Crushing
	d.Fire += damage.Fire
	d.Cold += damage.Cold
	d.Lighting += damage.Lighting
	d.Poison += damage.Poison
	d.Exhaustion += damage.Exhaustion
	d.ManaDrain += damage.ManaDrain
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
	d.Exhaustion -= damage.Exhaustion
	d.ManaDrain -= damage.ManaDrain
	d.Bleeding -= damage.Bleeding
	d.Fear -= damage.Fear
	d.Curse -= damage.Curse
}

func (d *Damage) HasPhysicalEffect() bool {
	return d.Stabbing != 0 ||
		d.Cutting != 0 ||
		d.Crushing != 0 ||
		d.Fire != 0 ||
		d.Cold != 0 ||
		d.Lighting != 0
}

func (d *Damage) HasEffect() bool {
	return d.HasPhysicalEffect() ||
		d.Poison != 0 ||
		d.Exhaustion != 0 ||
		d.ManaDrain != 0 ||
		d.Bleeding != 0 ||
		d.Fear != 0 ||
		d.Curse != 0
}

func (d *Damage) Enchance(attributes UnitAttributes, damage Damage) {
	d.Stabbing = util.AccumulateIfNotZeros(d.Stabbing, attributes.Strength)
	d.Cutting = util.AccumulateIfNotZeros(d.Cutting, attributes.Strength)
	d.Crushing = util.AccumulateIfNotZeros(d.Crushing, attributes.Strength)
	d.Fire = util.AccumulateIfNotZeros(d.Fire, attributes.Intelligence)
	d.Cold = util.AccumulateIfNotZeros(d.Cold, attributes.Intelligence)
	d.Lighting = util.AccumulateIfNotZeros(d.Lighting, attributes.Intelligence)
	d.Exhaustion = util.AccumulateIfNotZeros(d.Exhaustion, attributes.Strength)
	d.ManaDrain = util.AccumulateIfNotZeros(d.ManaDrain, attributes.Intelligence)
	d.Bleeding = util.AccumulateIfNotZeros(d.Bleeding, attributes.Strength)

	d.Stabbing = util.AccumulateIfNotZeros(d.Stabbing, damage.Stabbing)
	d.Cutting = util.AccumulateIfNotZeros(d.Cutting, damage.Cutting)
	d.Crushing = util.AccumulateIfNotZeros(d.Crushing, damage.Crushing)
	d.Fire = util.AccumulateIfNotZeros(d.Fire, damage.Fire)
	d.Cold = util.AccumulateIfNotZeros(d.Cold, damage.Cold)
	d.Lighting = util.AccumulateIfNotZeros(d.Lighting, damage.Lighting)
	d.Poison = util.AccumulateIfNotZeros(d.Poison, damage.Poison)
	d.Exhaustion = util.AccumulateIfNotZeros(d.Exhaustion, damage.Exhaustion)
	d.ManaDrain = util.AccumulateIfNotZeros(d.ManaDrain, damage.ManaDrain)
	d.Bleeding = util.AccumulateIfNotZeros(d.Bleeding, damage.Bleeding)
	d.Fear = util.AccumulateIfNotZeros(d.Fear, damage.Fear)
	d.Curse = util.AccumulateIfNotZeros(d.Curse, damage.Curse)
}

func (d *Damage) Multiply(factor float32) {
	d.Stabbing = util.MultiplyIfNotZeros(d.Stabbing, factor)
	d.Cutting = util.MultiplyIfNotZeros(d.Cutting, factor)
	d.Crushing = util.MultiplyIfNotZeros(d.Crushing, factor)
	d.Fire = util.MultiplyIfNotZeros(d.Fire, factor)
	d.Cold = util.MultiplyIfNotZeros(d.Cold, factor)
	d.Lighting = util.MultiplyIfNotZeros(d.Lighting, factor)
	d.Poison = util.MultiplyIfNotZeros(d.Poison, factor)
	d.Exhaustion = util.MultiplyIfNotZeros(d.Exhaustion, factor)
	d.ManaDrain = util.MultiplyIfNotZeros(d.ManaDrain, factor)
	d.Bleeding = util.MultiplyIfNotZeros(d.Bleeding, factor)
	d.Fear = util.MultiplyIfNotZeros(d.Fear, factor)
	d.Curse = util.MultiplyIfNotZeros(d.Curse, factor)
}

func (d *Damage) Normalize() {
	d.Stabbing = util.Max(d.Stabbing, 0)
	d.Cutting = util.Max(d.Cutting, 0)
	d.Crushing = util.Max(d.Crushing, 0)
	d.Fire = util.Max(d.Fire, 0)
	d.Cold = util.Max(d.Cold, 0)
	d.Lighting = util.Max(d.Lighting, 0)
	d.Poison = util.Max(d.Poison, 0)
	d.Exhaustion = util.Max(d.Exhaustion, 0)
	d.ManaDrain = util.Max(d.ManaDrain, 0)
	d.Bleeding = util.Max(d.Bleeding, 0)
	d.Fear = util.Max(d.Fear, 0)
	d.Curse = util.Max(d.Curse, 0)
}

func (d *Damage) Apply(state *UnitState) {
	state.Health -= d.Stabbing + d.Cutting + d.Crushing + d.Fire + d.Cold + d.Lighting + d.Poison + d.Bleeding
	state.Stamina -= d.Exhaustion
	state.Mana -= d.ManaDrain
	state.Curse += d.Curse
	state.Fear += d.Fear

	state.Health = util.Max(state.Health, 0)
	state.Stamina = util.Max(state.Stamina, 0)
	state.Mana = util.Max(state.Mana, 0)
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
	if d.Exhaustion != 0 {
		props = append(props, fmt.Sprintf("exhaustion: %g", d.Exhaustion))
	}
	if d.ManaDrain != 0 {
		props = append(props, fmt.Sprintf("mana drain: %g", d.ManaDrain))
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
