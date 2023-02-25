package domain

import "jrpg-gang/util"

type UnitResistance struct {
	Damage `bson:",inline"`
}

func (r *UnitResistance) Accumulate(resistance UnitResistance) {
	r.Damage.Accumulate(resistance.Damage)
}

func (r *UnitResistance) AccumulatePhysical(value float32) {
	r.Stabbing += value
	r.Cutting += value
	r.Crushing += value
	r.Fire += value
	r.Cold += value
	r.Lighting += value
}

func (r *UnitResistance) PhysicalAbsorption(damage Damage) float32 {
	return util.Min(r.Stabbing, damage.Stabbing) +
		util.Min(r.Cutting, damage.Cutting) +
		util.Min(r.Crushing, damage.Crushing) +
		util.Min(r.Fire, damage.Fire) +
		util.Min(r.Cold, damage.Cold) +
		util.Min(r.Lighting, damage.Lighting)
}
