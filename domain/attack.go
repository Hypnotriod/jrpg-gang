package domain

import (
	"math/rand"
)

func (u *Unit) ApplyDamage(damage Damage) {
	// todo: implement
}

func (u *Unit) Attack(target *Unit, impact []DamageImpact) (*Damage, bool) {
	var damage *Damage = &Damage{}
	var success bool = false
	for _, i := range impact {
		// todo: find better formula?
		chance := i.Chance + u.Stats.Attributes.Luck
		rnd := rand.Float32() * 100
		if chance <= rnd {
			continue
		}
		success = true
		damage.Accumulate(&i.Damage)
		switch i.Type {
		case ImpactTypeTemporal, ImpactTypePermanent:
			u.Impact = append(u.Impact, i)
		}
	}

	resistance := target.TotalResistance()
	damage.Reduce(&resistance.Damage)
	damage.Normalize()
	target.ApplyDamage(*damage)
	return damage, success
}
