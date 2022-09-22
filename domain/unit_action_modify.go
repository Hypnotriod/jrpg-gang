package domain

func (u *Unit) Modify(target *Unit, modification []UnitModificationImpact) ([]UnitRecovery, []UnitModificationImpact) {
	instantRecovery := []UnitRecovery{}
	temporalModification := []UnitModificationImpact{}
	for _, ench := range modification {
		if ench.Chance != 0 && !u.CheckRandomChance(u.CalculateModificationChance(ench)) {
			break
		}
		if ench.Duration != 0 {
			target.Modification = append(target.Modification, ench)
			temporalModification = append(temporalModification, ench)
		} else {
			target.ApplyRecovery(ench.Recovery)
			instantRecovery = append(instantRecovery, ench.Recovery)
		}
	}
	return instantRecovery, temporalModification
}

func (u *Unit) ApplyRecovery(recovery UnitRecovery) {
	modification := u.TotalModification()
	attributes := modification.BaseAttributes
	attributes.Accumulate(u.Stats.BaseAttributes)
	u.State.Accumulate(recovery.UnitState)
	u.State.Saturate(attributes)
	u.State.Normalize()
}
