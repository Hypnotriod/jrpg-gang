package domain

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
	r.Lightning += value
}

func (r *UnitResistance) PhysicalAbsorption(damage Damage) float32 {
	return min(r.Stabbing, damage.Stabbing) +
		min(r.Cutting, damage.Cutting) +
		min(r.Crushing, damage.Crushing) +
		min(r.Fire, damage.Fire) +
		min(r.Cold, damage.Cold) +
		min(r.Lightning, damage.Lightning)
}

func (r UnitResistance) IsZero() bool {
	return r.Damage.IsZero()
}
