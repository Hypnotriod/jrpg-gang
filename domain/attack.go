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

func (u *Unit) Attack(target *Unit, impact []DamageImpact) (Damage, []DamageImpact, bool) {
	var damage Damage = Damage{}
	var tempImpact []DamageImpact = []DamageImpact{}
	var success bool = false
	for _, imp := range impact {
		// todo: find better formula?
		chance := imp.Chance + u.Stats.Attributes.Luck - u.State.Curse
		if !util.CheckRandomChance(chance) {
			continue
		}
		if imp.Duration != 0 {
			imp.Chance = 0
			if target.ApplyTemporalImpact(imp) {
				tempImpact = append(tempImpact, imp)
				success = true
			}
		} else {
			damage.Accumulate(imp.Damage)
			success = true
		}
	}
	if success {
		damage.Enchance(u.Stats.Attributes)
		damage = target.ApplyPermanentDamage(damage)
	}
	return damage, tempImpact, success
}
