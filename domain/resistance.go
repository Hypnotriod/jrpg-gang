package domain

type UnitResistance struct {
	Damage
}

func (r *UnitResistance) Accumulate(resistance UnitResistance) {
	r.Damage.Accumulate(resistance.Damage)
}
