package domain

import "jrpg-gang/util"

func (u *Unit) CalculateEnchancementChance(enhancement UnitEnhancementImpact) float32 {
	chance := (u.TotalIntelligence() - u.State.Curse) + enhancement.Chance
	return util.MaxFloat32(chance, util.MINIMUM_CHANCE)
}

func (u *Unit) ApplyRecovery(recovery UnitState) {
	enhancement := u.TotalEnhancement()
	attributes := enhancement.BaseAttributes
	attributes.Accumulate(u.Stats.BaseAttributes)
	u.State.Accumulate(recovery)
	u.State.Normalize(attributes)
}

func (u *Unit) Enhance(target *Unit, enhancement []UnitEnhancementImpact) ([]UnitState, []UnitEnhancementImpact) {
	instantRecovery := []UnitState{}
	temporalEnhancement := []UnitEnhancementImpact{}
	for _, ench := range enhancement {
		if ench.Chance != 0 && !util.CheckRandomChance(u.CalculateEnchancementChance(ench)) {
			break
		}
		if ench.Duration != 0 {
			target.Enhancement = append(target.Enhancement, ench)
			temporalEnhancement = append(temporalEnhancement, ench)
		} else {
			target.ApplyRecovery(ench.Recovery)
			instantRecovery = append(instantRecovery, ench.Recovery)
		}
	}
	return instantRecovery, temporalEnhancement
}
