package domain

import (
	"jrpg-gang/util"
)

func (u *Unit) ApplyDamage(damage Damage) {
	resistance := u.TotalResistance()
	damage.Reduce(resistance.Damage)
	damage.Normalize()
	damage.Apply(&u.State)
}

func (u *Unit) Attack(target *Unit, impact []DamageImpact) (Damage, bool) {
	var damage Damage = Damage{}
	var success bool = false
	for _, imp := range impact {
		// todo: find better formula?
		chance := imp.Chance + u.Stats.Attributes.Luck - u.State.Curse
		if !util.CheckRandomChance(chance) {
			continue
		}
		success = true
		damage.Accumulate(imp.Damage)
		if imp.Duration != 0 {
			u.Impact = append(u.Impact, imp)
		}
	}
	if success {
		damage.Enchance(u.Stats.Attributes)
		target.ApplyDamage(damage)
	}
	return damage, success
}
