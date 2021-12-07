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

func (u *Unit) CalculateCritilalAttackChance(target *Unit) float32 {
	chance := (u.TotalLuck() - u.State.Curse) - (target.TotalLuck() - target.State.Curse)
	return util.MaxFloat32(chance, util.MINIMUM_CHANCE)
}

func (u *Unit) CalculateAttackChance(target *Unit, impact DamageImpact) float32 {
	chance := (u.TotalAgility() - u.State.Curse) - (target.TotalAgility() - target.State.Curse) + impact.Chance
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
