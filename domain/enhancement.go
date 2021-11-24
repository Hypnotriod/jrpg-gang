package domain

type UnitEnhancement struct {
	UnitBaseAttributes
	UnitAttributes
	UnitResistance
}

func (u *UnitEnhancement) Accumulate(enhancement UnitEnhancement) {
	u.UnitBaseAttributes.Health += enhancement.UnitBaseAttributes.Health
	u.UnitBaseAttributes.Mana += enhancement.UnitBaseAttributes.Mana
	u.UnitBaseAttributes.Stamina += enhancement.UnitBaseAttributes.Stamina

	u.UnitAttributes.Strength += enhancement.UnitAttributes.Strength
	u.UnitAttributes.Physique += enhancement.UnitAttributes.Physique
	u.UnitAttributes.Dexterity += enhancement.UnitAttributes.Dexterity
	u.UnitAttributes.Endurance += enhancement.UnitAttributes.Endurance
	u.UnitAttributes.Intelligence += enhancement.UnitAttributes.Intelligence
	u.UnitAttributes.Luck += enhancement.UnitAttributes.Luck
}

func (u *Unit) TotalLuck(checkEquipment bool) float32 {
	var luck float32 = u.Stats.Attributes.Luck - u.State.Curse
	for _, e := range u.Enhancement {
		luck += e.Luck
	}
	if !checkEquipment {
		return luck
	}
	for _, item := range u.Items {
		equipment, ok := AsEquipment(item)
		if !ok || !equipment.Equipped {
			continue
		}
		for _, e := range equipment.Enhancement {
			luck += e.Luck
		}
	}
	return luck
}

func (u *Unit) TotalEnhancement(checkEquipment bool) *UnitEnhancement {
	var enhancement *UnitEnhancement = &UnitEnhancement{}
	for _, e := range u.Enhancement {
		enhancement.Accumulate(e.UnitEnhancement)
	}
	if !checkEquipment {
		return enhancement
	}
	for _, item := range u.Items {
		equipment, ok := AsEquipment(item)
		if !ok || !equipment.Equipped {
			continue
		}
		for _, e := range equipment.Enhancement {
			enhancement.Accumulate(e)
		}
	}
	return enhancement
}
