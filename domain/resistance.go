package domain

type UnitResistance struct {
	Damage
}

func (r *UnitResistance) Accumulate(resistance *UnitResistance) {
	r.Stabbing += resistance.Stabbing
	r.Cutting += resistance.Cutting
	r.Crushing += resistance.Crushing
	r.Fire += resistance.Fire
	r.Cold += resistance.Cold
	r.Lighting += resistance.Lighting
	r.Fear += resistance.Fear
	r.Poison += resistance.Poison
	r.Curse += resistance.Curse
	r.Stunning += resistance.Stunning
}

func AccumulateResistanceByEquipment(resistance *UnitResistance, item interface{}, checkEquipped bool) *UnitResistance {
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
	for _, item := range u.Items {
		AccumulateResistanceByEquipment(resistance, item, true)
	}
	return resistance
}
