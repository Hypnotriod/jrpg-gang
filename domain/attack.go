package domain

import (
	"jrpg-gang/util"
)

func (u *Unit) ApplyInstantDamage(damage Damage) Damage {
	resistance := u.TotalResistance()
	damage.Reduce(resistance.Damage)
	damage.Normalize()
	damage.Apply(&u.State)
	return damage
}

func (u *Unit) ApplyTemporalImpact(impact DamageImpact) bool {
	resistance := u.TotalResistance()
	impact.Reduce(resistance.Damage)
	impact.Normalize()
	if impact.HasEffect() {
		u.Impact = append(u.Impact, impact)
		return true
	}
	return false
}

func (u *Unit) CalculateCritilalAttackChance() float32 {
	return u.TotalLuck(true)
}

func (u *Unit) CalculateAttackChance(target *Unit, impact DamageImpact) float32 {
	chance := u.TotalAgility(true) - target.TotalAgility(true) + impact.Chance
	return util.MaxFloat32(chance, MINIMAL_CHANCE)
}

func (u *Unit) Attack(target *Unit, impact []DamageImpact) (Damage, []DamageImpact) {
	var instantDamage Damage = Damage{}
	var temporalImpact []DamageImpact = []DamageImpact{}
	for _, imp := range impact {
		chance := u.CalculateAttackChance(target, imp)
		if !util.CheckRandomChance(chance) {
			continue
		}
		if imp.Duration != 0 {
			imp.Chance = 0
			if target.ApplyTemporalImpact(imp) {
				temporalImpact = append(temporalImpact, imp)
			}
		} else {
			instantDamage.Accumulate(imp.Damage)
		}
	}
	if instantDamage.HasEffect() {
		instantDamage.Enchance(u.Stats.Attributes)
		criticalChance := u.CalculateCritilalAttackChance()
		if util.CheckRandomChance(criticalChance) {
			instantDamage.Multiply(CRITICAL_FACTOR)
		}
		instantDamage = target.ApplyInstantDamage(instantDamage)
	}
	return instantDamage, temporalImpact
}
