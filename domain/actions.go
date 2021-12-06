package domain

import "jrpg-gang/util"

func (u *Unit) ApplyDamage(damage Damage) Damage {
	resistance := u.TotalEnhancement().Resistance
	resistance.Accumulate(u.Stats.Resistance)
	damage.Reduce(resistance.Damage)
	damage.Normalize()
	if damage.HasEffect() {
		damage.Apply(&u.State)
	}
	return damage
}

func (u *Unit) AccumulateImpact(impact DamageImpact) DamageImpact {
	resistance := u.TotalEnhancement().Resistance
	resistance.Accumulate(u.Stats.Resistance)
	impact.Reduce(resistance.Damage)
	impact.Normalize()
	if impact.HasEffect() {
		u.Impact = append(u.Impact, impact)
	}
	return impact
}

func (u *Unit) ApplyRecovery(recovery UnitState) {
	enhancement := u.TotalEnhancement()
	attributes := enhancement.BaseAttributes
	attributes.Accumulate(u.Stats.BaseAttributes)
	u.State.Accumulate(recovery)
	u.State.Normalize(attributes)
}

func (u *Unit) CalculateCritilalAttackChance(target *Unit) float32 {
	chance := (u.TotalLuck() - u.State.Curse) - (target.TotalLuck() - target.State.Curse)
	return util.MaxFloat32(chance, util.MINIMUM_CHANCE)
}

func (u *Unit) CalculateAttackChance(target *Unit, impact DamageImpact) float32 {
	chance := (u.TotalAgility() - u.State.Curse) - (target.TotalAgility() - target.State.Curse) + impact.Chance
	return util.MaxFloat32(chance, util.MINIMUM_CHANCE)
}

func (u *Unit) CalculateEnchancementChance(enhancement UnitEnhancementImpact) float32 {
	chance := (u.TotalIntelligence() - u.State.Curse) + enhancement.Chance
	return util.MaxFloat32(chance, util.MINIMUM_CHANCE)
}

func (u *Unit) Attack(target *Unit, impact []DamageImpact) ([]Damage, []DamageImpact) {
	instantDamage := []Damage{}
	temporalImpact := []DamageImpact{}
	totalEnhancement := u.TotalEnhancement()
	totalEnhancement.Attributes.Accumulate(u.Stats.Attributes)
	for _, imp := range impact {
		if !util.CheckRandomChance(u.CalculateAttackChance(target, imp)) {
			break
		}
		imp.Enchance(totalEnhancement.Attributes, totalEnhancement.Damage)
		if util.CheckRandomChance(u.CalculateCritilalAttackChance(target)) {
			imp.Damage.Multiply(CRITICAL_DAMAGE_FACTOR)
			imp.Damage.IsCritical = true
		}
		if imp.Duration != 0 {
			tmpImp := target.AccumulateImpact(imp)
			tmpImp.Chance = 0
			temporalImpact = append(temporalImpact, tmpImp)
		} else {
			instDmg := target.ApplyDamage(imp.Damage)
			instantDamage = append(instantDamage, instDmg)
		}
	}
	return instantDamage, temporalImpact
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
