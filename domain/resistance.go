package domain

type UnitResistance struct {
	Damage
}

func (r *UnitResistance) Accumulate(resistance *UnitResistance) {
	r.Damage.Accumulate(resistance.Damage)
}

func (u *Unit) ItemsResistance(checkEquipped bool) *UnitResistance {
	var resistance *UnitResistance = &UnitResistance{}
	for _, item := range u.Items {
		resistance.Accumulate(u.ItemResistance(item, checkEquipped))
	}
	return resistance
}

func (u *Unit) ItemResistance(item interface{}, checkEquipped bool) *UnitResistance {
	var resistance *UnitResistance = &UnitResistance{}
	equipment, ok := AsEquipment(item)
	if !ok || checkEquipped && !equipment.Equipped {
		return resistance
	}
	for _, enhancement := range equipment.Enhancement {
		resistance.Accumulate(&enhancement.UnitResistance)
	}
	return resistance
}

func (u *Unit) TotalResistance() *UnitResistance {
	var resistance *UnitResistance = &UnitResistance{}
	resistance.Accumulate(&u.Stats.Resistance)
	resistance.Accumulate(u.ItemsResistance(true))
	resistance.Normalize()
	return resistance
}
