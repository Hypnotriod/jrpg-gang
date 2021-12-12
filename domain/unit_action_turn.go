package domain

func (u *Unit) ApplyDamageOnNextTurn() Damage {
	var damage Damage
	for i := range u.Damage {
		damage.Accumulate(u.Damage[i].Damage)
		if u.Damage[i].Duration > 0 {
			u.Damage[i].Duration--
		}
	}
	damage.Apply(&u.State)
	u.FilterDamage()
	return damage
}

func (u *Unit) FilterDamage() {
	var filteredDamage []DamageImpact
	for _, damage := range u.Damage {
		if damage.Duration != 0 {
			filteredDamage = append(filteredDamage, damage)
		}
	}
	u.Damage = filteredDamage
}

func (u *Unit) ApplyRecoverylOnNextTurn() {
	modification := u.TotalModification()
	recovery := modification.Recovery
	attributes := modification.Attributes
	baseAttributes := modification.BaseAttributes
	attributes.Accumulate(u.Stats.Attributes)
	baseAttributes.Accumulate(u.Stats.BaseAttributes)
	attributes.Normalize()
	baseAttributes.Normalize()
	recovery.Normalize()
	u.State.Stamina += attributes.Endurance
	u.State.Accumulate(recovery)
	u.State.NormalizeWithLimit(baseAttributes)
}

func (u *Unit) ReduceModificationOnNextTurn() {
	for i := range u.Modification {
		if u.Modification[i].Duration > 0 {
			u.Modification[i].Duration--
		}
	}
	u.FilterModification()
}

func (u *Unit) FilterModification() {
	var filteredModification []UnitModificationImpact
	for _, modification := range u.Modification {
		if modification.Duration != 0 {
			filteredModification = append(filteredModification, modification)
		}
	}
	u.Modification = filteredModification
}
