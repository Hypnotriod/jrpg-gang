package domain

func (u *Unit) Modify(target *Unit, modification []UnitModificationImpact) ([]UnitRecovery, []UnitModificationImpact) {
	instantRecovery := []UnitRecovery{}
	temporalModification := []UnitModificationImpact{}
	intelligence := u.TotalIntelligence()
	for _, imp := range modification {
		if imp.Chance != 0 && !u.CheckRandomChance(u.CalculateModificationChance(imp)) {
			break
		}
		imp.EnchanceAll(u.PickDeviation(imp.Deviation))
		imp.MultiplyAll(1 + intelligence*INTELLIGENCE_MODIFICATION_FACTOR)
		imp.Chance = 0
		if imp.Duration != 0 {
			target.Modification = append(target.Modification, imp)
			temporalModification = append(temporalModification, imp)
		} else {
			target.ApplyRecovery(imp.Recovery)
			instantRecovery = append(instantRecovery, imp.Recovery)
		}
	}
	return instantRecovery, temporalModification
}

func (u *Unit) ApplyRecovery(recovery UnitRecovery) {
	modification := u.TotalModification()
	attributes := modification.BaseAttributes
	attributes.Accumulate(u.Stats.BaseAttributes)
	u.ReduceDamageImpact(recovery.Damage)
	u.State.Accumulate(recovery.UnitState)
	u.State.Saturate(attributes)
	u.State.Normalize()
}
