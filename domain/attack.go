package domain

import (
	"jrpg-gang/util"
)

func (u *Unit) ApplyInstantDamage(damage Damage) Damage {
	totalEnhancement := u.TotalEnhancement(true)
	totalEnhancement.Resistance.Accumulate(u.Stats.Resistance)
	damage.Reduce(totalEnhancement.Resistance.Damage)
	damage.Normalize()
	damage.Apply(&u.State)
	return damage
}

func (u *Unit) ApplyTemporalImpact(impact DamageImpact) DamageImpact {
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
	return u.TotalLuck(true)
}

func (u *Unit) CalculateAttackChance(target *Unit, impact DamageImpact) float32 {
	chance := u.TotalAgility(true) - target.TotalAgility(true) + impact.Chance
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
			if tmpImp := target.ApplyTemporalImpact(imp); tmpImp.HasEffect() {
				temporalImpact = append(temporalImpact, imp)
			}
		} else {
			if instDmg := target.ApplyInstantDamage(imp.Damage); instDmg.HasEffect() {
				instantDamage = append(instantDamage, instDmg)
			}
		}
	}
	return instantDamage, temporalImpact
}
