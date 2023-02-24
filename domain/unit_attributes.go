package domain

import (
	"jrpg-gang/util"
)

type UnitAttributes struct {
	Strength     float32 `json:"strength" bson:"strength"`         // enhances (stabbing, cutting, crushing, bleeding) damage
	Physique     float32 `json:"physique" bson:"physique"`         // affects stun chance
	Agility      float32 `json:"agility" bson:"agility"`           // affects attack chance, dodge chance
	Endurance    float32 `json:"endurance" bson:"endurance"`       // stamina recovery
	Intelligence float32 `json:"intelligence" bson:"intelligence"` // enhances (fire, cold, lighting, exhaustion, manaDrain, fear, curse) damage, adds 1% to all modification points
	Initiative   float32 `json:"initiative" bson:"initiative"`     // affects turn order
	Luck         float32 `json:"luck" bson:"luck"`                 // affects critical chance
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

func (a *UnitAttributes) Normalize() {
	a.Strength = util.Max(a.Strength, 0)
	a.Physique = util.Max(a.Physique, 0)
	a.Agility = util.Max(a.Agility, 0)
	a.Endurance = util.Max(a.Endurance, 0)
	a.Intelligence = util.Max(a.Intelligence, 0)
	a.Initiative = util.Max(a.Initiative, 0)
	a.Luck = util.Max(a.Luck, 0)
}

func (a *UnitAttributes) MultiplyAll(factor float32) {
	a.Strength = util.MultiplyWithRounding(a.Strength, factor)
	a.Physique = util.MultiplyWithRounding(a.Physique, factor)
	a.Agility = util.MultiplyWithRounding(a.Agility, factor)
	a.Endurance = util.MultiplyWithRounding(a.Endurance, factor)
	a.Intelligence = util.MultiplyWithRounding(a.Intelligence, factor)
	a.Initiative = util.MultiplyWithRounding(a.Initiative, factor)
	a.Luck = util.MultiplyWithRounding(a.Luck, factor)
}

func (a *UnitAttributes) CheckRequirements(requirements UnitAttributes) bool {
	return a.Strength >= requirements.Strength &&
		a.Physique >= requirements.Physique &&
		a.Agility >= requirements.Agility &&
		a.Endurance >= requirements.Endurance &&
		a.Intelligence >= requirements.Intelligence &&
		a.Initiative >= requirements.Initiative &&
		a.Luck >= requirements.Luck
}
