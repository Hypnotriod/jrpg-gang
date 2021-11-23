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
	for _, i := range impact {
		// todo: find better formula?
		chance := i.Chance + u.Stats.Attributes.Luck - u.State.Curse
		if !util.CheckRandomChance(chance) {
			continue
		}
		success = true
		damage.Accumulate(i.Damage)
		switch i.Type {
		case ImpactTypeTemporal, ImpactTypePermanent:
			u.Impact = append(u.Impact, i)
		}
	}
	if success {
		damage.Enchance(u.Stats.Attributes)
		target.ApplyDamage(damage)
	}
	return damage, success
}
