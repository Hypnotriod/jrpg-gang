package domain

import "jrpg-gang/util"

func (u *Unit) ApplyDamage(damage Damage) Damage {
	totalEnhancement := u.TotalEnhancement(true)
	totalEnhancement.Resistance.Accumulate(u.Stats.Resistance)
	damage.Reduce(totalEnhancement.Resistance.Damage)
	damage.Normalize()
	damage.Apply(&u.State)
	return damage
}

func (u *Unit) AccumulateImpact(impact DamageImpact) DamageImpact {
	totalEnhancement := u.TotalEnhancement(true)
	totalEnhancement.Resistance.Accumulate(u.Stats.Resistance)
	impact.Reduce(totalEnhancement.Resistance.Damage)
	impact.Normalize()
	if impact.HasEffect() {
		u.Impact = append(u.Impact, impact)
	}
	return impact
}

func (u *Unit) CalculateCritilalAttackChance() float32 {
	return u.TotalLuck(true) - u.State.Curse
}

func (u *Unit) CalculateAttackChance(target *Unit, impact DamageImpact) float32 {
	chance := (u.TotalAgility(true) - u.State.Curse) - (target.TotalAgility(true) - target.State.Curse) + impact.Chance
	return util.MaxFloat32(chance, MINIMAL_CHANCE)
}

func (u *Unit) Attack(target *Unit, impact []DamageImpact) ([]Damage, []DamageImpact) {
	instantDamage := []Damage{}
	temporalImpact := []DamageImpact{}
	totalEnhancement := u.TotalEnhancement(true)
	totalEnhancement.Attributes.Accumulate(u.Stats.Attributes)
	for _, imp := range impact {
		if !util.CheckRandomChance(u.CalculateAttackChance(target, imp)) {
			continue
		}
		imp.Enchance(totalEnhancement.Attributes, totalEnhancement.Damage)
		if util.CheckRandomChance(u.CalculateCritilalAttackChance()) {
			imp.Damage.Multiply(CRITICAL_FACTOR)
			imp.Damage.IsCritical = true
		}
		if imp.Duration != 0 {
			imp.Chance = 0
			if tmpImp := target.AccumulateImpact(imp); tmpImp.HasEffect() {
				temporalImpact = append(temporalImpact, imp)
			}
		} else {
			if instDmg := target.ApplyDamage(imp.Damage); instDmg.HasEffect() {
				instantDamage = append(instantDamage, instDmg)
			}
		}
	}
	return instantDamage, temporalImpact
}

func (u *Unit) ApplyImpactOnNextTurn() Damage {
	var damage Damage
	for _, impact := range u.Impact {
		damage.Accumulate(impact.Damage)
		if impact.Duration > 0 {
			impact.Duration--
		}
	}
	damage.Apply(&u.State)
	u.FilterImpact()
	return damage
}

func (u *Unit) FilterImpact() {
	var filteredImpact []DamageImpact
	for _, impact := range u.Impact {
		if impact.Duration == 0 {
			filteredImpact = append(filteredImpact, impact)
		}
	}
	u.Impact = filteredImpact
}

func (u *Unit) ReduceEnhancementOnNextTurn() {
	for _, enhancement := range u.Enhancement {
		if enhancement.Duration > 0 {
			enhancement.Duration--
		}
	}
	u.FilterEnhancement()
}

func (u *Unit) FilterEnhancement() {
	var filteredEnhancement []UnitEnhancementImpact
	for _, enhancement := range u.Enhancement {
		if enhancement.Duration == 0 {
			filteredEnhancement = append(filteredEnhancement, enhancement)
		}
	}
	u.Enhancement = filteredEnhancement
}
