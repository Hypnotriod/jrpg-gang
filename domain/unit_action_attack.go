package domain

import "jrpg-gang/util"

func (u *Unit) Attack(target *Unit, damage []DamageImpact) ([]Damage, []DamageImpact) {
	wasStunned := target.State.IsStunned
	instantDamage := []Damage{}
	temporalDamage := []DamageImpact{}
	totalModification := u.TotalModification()
	totalModification.Attributes.Accumulate(u.Stats.Attributes)
	totalModification.Attributes.Normalize()
	for _, imp := range damage {
		if !u.CheckRandomChance(u.calculateAttackChance(target, imp)) {
			break
		}
		imp.Enchance(totalModification.Attributes, totalModification.Damage)
		imp.Normalize()
		if wasStunned || u.CheckRandomChance(u.calculateCritilalAttackChance(target)) {
			imp.Damage.Multiply(CRITICAL_DAMAGE_FACTOR)
			imp.Damage.IsCritical = true
		}
		if imp.Duration != 0 {
			tmpImp := target.accumulateDamageImpact(imp)
			tmpImp.Chance = 0
			temporalDamage = append(temporalDamage, tmpImp)
		} else {
			instDmg := target.applyDamage(imp.Damage)
			if !wasStunned && u.CheckRandomChance(target.calculateStunChance(u, imp)) {
				target.State.IsStunned = true
				instDmg.IsStunned = true
			}
			instantDamage = append(instantDamage, instDmg)
		}
	}
	return instantDamage, temporalDamage
}

func (u *Unit) applyDamage(damage Damage) Damage {
	if damage.HasPhysicalEffect() {
		u.Inventory.IncreaseArmorWearout(true)
	}
	resistance := u.TotalModification().Resistance
	resistance.Accumulate(u.Stats.Resistance)
	resistance.Normalize()
	damage.Reduce(resistance.Damage)
	damage.Normalize()
	if damage.HasEffect() {
		damage.Apply(&u.State)
	}
	return damage
}

func (u *Unit) accumulateDamageImpact(damage DamageImpact) DamageImpact {
	resistance := u.TotalModification().Resistance
	resistance.Accumulate(u.Stats.Resistance)
	resistance.Normalize()
	damage.Reduce(resistance.Damage)
	damage.Normalize()
	if damage.HasEffect() {
		u.Damage = append(u.Damage, damage)
	}
	return damage
}

func (u *Unit) calculateStunChance(target *Unit, damage DamageImpact) float32 {
	chance := (damage.PhysicalDamage() - target.State.Curse) - (u.TotalPhysique() - u.State.Curse)
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) calculateCritilalAttackChance(target *Unit) float32 {
	chance := (u.TotalLuck() - u.State.Curse) - (target.TotalLuck() - target.State.Curse)
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) calculateAttackChance(target *Unit, damage DamageImpact) float32 {
	chance := (u.TotalAgility() - u.State.Curse) - (target.TotalAgility() - target.State.Curse) + damage.Chance
	return util.Max(chance, MINIMUM_CHANCE)
}
