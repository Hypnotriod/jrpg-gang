package domain

import (
	"jrpg-gang/util"
)

type Damage struct {
	Stabbing   float32 `json:"stabbing,omitempty"`   // affects health
	Cutting    float32 `json:"cutting,omitempty"`    // affects health
	Crushing   float32 `json:"crushing,omitempty"`   // affects health
	Fire       float32 `json:"fire,omitempty"`       // affects health
	Cold       float32 `json:"cold,omitempty"`       // affects health
	Lighting   float32 `json:"lighting,omitempty"`   // affects health
	Poison     float32 `json:"poison,omitempty"`     // affects health
	Exhaustion float32 `json:"exhaustion,omitempty"` // affects stamina
	ManaDrain  float32 `json:"manaDrain,omitempty"`  // affects mana
	Bleeding   float32 `json:"bleeding,omitempty"`   // affects health
	Morale     float32 `json:"morale,omitempty"`     // affects fear
	Fortune    float32 `json:"fortune,omitempty"`    // affects curse
	IsCritical bool    `json:"isCritical,omitempty"` // critical damage flag
	WithStun   bool    `json:"withStun,omitempty"`   // stun flag
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
	d.Morale += damage.Morale
	d.Fortune += damage.Fortune
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
	d.Morale -= damage.Morale
	d.Fortune -= damage.Fortune
}

func (d *Damage) PhysicalDamage() float32 {
	return d.Stabbing + d.Cutting + d.Crushing + d.Fire + d.Cold + d.Lighting
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
		d.Morale != 0 ||
		d.Fortune != 0
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
	d.Morale = util.AccumulateIfNotZeros(d.Morale, attributes.Intelligence)
	d.Fortune = util.AccumulateIfNotZeros(d.Fortune, attributes.Intelligence)

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
	d.Morale = util.AccumulateIfNotZeros(d.Morale, damage.Morale)
	d.Fortune = util.AccumulateIfNotZeros(d.Fortune, damage.Fortune)
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
	d.Morale = util.MultiplyIfNotZeros(d.Morale, factor)
	d.Fortune = util.MultiplyIfNotZeros(d.Fortune, factor)
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
	d.Morale = util.Max(d.Morale, 0)
	d.Fortune = util.Max(d.Fortune, 0)
}

func (d *Damage) Apply(state *UnitState) {
	state.Health -= d.Stabbing + d.Cutting + d.Crushing + d.Fire + d.Cold + d.Lighting + d.Poison + d.Bleeding
	state.Stamina -= d.Exhaustion
	state.Mana -= d.ManaDrain
	state.Curse += d.Fortune
	state.Fear += d.Morale

	state.Health = util.Max(state.Health, 0)
	state.Stamina = util.Max(state.Stamina, 0)
	state.Mana = util.Max(state.Mana, 0)
}
