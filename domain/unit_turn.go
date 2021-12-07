package domain

func (u *Unit) ApplyDamageImpactOnNextTurn() Damage {
	var damage Damage
	for i := range u.Impact {
		damage.Accumulate(u.Impact[i].Damage)
		if u.Impact[i].Duration > 0 {
			u.Impact[i].Duration--
		}
	}
	damage.Apply(&u.State)
	u.FilterImpact()
	return damage
}

func (u *Unit) FilterImpact() {
	var filteredImpact []DamageImpact
	for _, impact := range u.Impact {
		if impact.Duration != 0 {
			filteredImpact = append(filteredImpact, impact)
		}
	}
	u.Impact = filteredImpact
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
