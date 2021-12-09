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

func (u *Unit) ApplyRecoverylEnhancementOnNextTurn() {
	enhancement := u.TotalEnhancement()
	recovery := enhancement.Recovery
	attributes := enhancement.BaseAttributes
	attributes.Accumulate(u.Stats.BaseAttributes)
	u.State.Accumulate(recovery)
	u.State.Normalize(attributes)
}

func (u *Unit) ReduceEnhancementOnNextTurn() {
	for i := range u.Enhancement {
		if u.Enhancement[i].Duration > 0 {
			u.Enhancement[i].Duration--
		}
	}
	u.FilterEnhancement()
}

func (u *Unit) FilterEnhancement() {
	var filteredEnhancement []UnitEnhancementImpact
	for _, enhancement := range u.Enhancement {
		if enhancement.Duration != 0 {
			filteredEnhancement = append(filteredEnhancement, enhancement)
		}
	}
	u.Enhancement = filteredEnhancement
}
