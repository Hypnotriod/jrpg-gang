package domain

import "jrpg-gang/util"

func (u *Unit) CalculateModificationChance(modification UnitModificationImpact) float32 {
	chance := (u.TotalIntelligence() - u.State.Curse) + modification.Chance
	return util.MaxFloat32(chance, util.MINIMUM_CHANCE)
}

func (u *Unit) ApplyRecovery(recovery UnitState) {
	modification := u.TotalModification()
	attributes := modification.BaseAttributes
	attributes.Accumulate(u.Stats.BaseAttributes)
	u.State.Accumulate(recovery)
	u.State.Saturate(attributes)
	u.State.Normalize()
}

func (u *Unit) Modify(target *Unit, modification []UnitModificationImpact) ([]UnitState, []UnitModificationImpact) {
	instantRecovery := []UnitState{}
	temporalModification := []UnitModificationImpact{}
	for _, ench := range modification {
		if ench.Chance != 0 && !util.CheckRandomChance(u.CalculateModificationChance(ench)) {
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
