package domain

import (
	"jrpg-gang/util"
)

func (u *Unit) ApplyPermanentDamage(damage Damage) Damage {
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

func (u *Unit) CalculateAttackChance(target *Unit, impact DamageImpact) float32 {
	// todo: find better formula
	return u.TotalLuck(true) + impact.Chance
}

func (u *Unit) Attack(target *Unit, impact []DamageImpact) (Damage, []DamageImpact, bool) {
	var permanentDamage Damage = Damage{}
	var temporalImpact []DamageImpact = []DamageImpact{}
	var success bool = false
	for _, imp := range impact {
		chance := u.CalculateAttackChance(target, imp)
		if !util.CheckRandomChance(chance) {
			continue
		}
		if imp.Duration != 0 {
			imp.Chance = 0
			if target.ApplyTemporalImpact(imp) {
				temporalImpact = append(temporalImpact, imp)
				success = true
			}
		} else {
			permanentDamage.Accumulate(imp.Damage)
			success = true
		}
	}
	if success {
		permanentDamage.Enchance(u.Stats.Attributes)
		permanentDamage = target.ApplyPermanentDamage(permanentDamage)
	}
	return permanentDamage, temporalImpact, success
}
