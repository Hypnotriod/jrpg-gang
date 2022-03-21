package domain

import "jrpg-gang/util"

func (u *Unit) ApplyDamage(damage Damage) Damage {
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

func (u *Unit) AccumulateDamageImpact(damage DamageImpact) DamageImpact {
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

func (u *Unit) CalculateCritilalAttackChance(target *Unit) float32 {
	chance := (u.TotalLuck() - u.State.Curse) - (target.TotalLuck() - target.State.Curse)
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateAttackChance(target *Unit, damage DamageImpact) float32 {
	chance := (u.TotalAgility() - u.State.Curse) - (target.TotalAgility() - target.State.Curse) + damage.Chance
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) Attack(target *Unit, damage []DamageImpact) ([]Damage, []DamageImpact) {
	instantDamage := []Damage{}
	temporalDamage := []DamageImpact{}
	totalModification := u.TotalModification()
	totalModification.Attributes.Accumulate(u.Stats.Attributes)
	totalModification.Attributes.Normalize()
	for _, imp := range damage {
		if !u.CheckRandomChance(u.CalculateAttackChance(target, imp)) {
			break
		}
		imp.Enchance(totalModification.Attributes, totalModification.Damage)
		imp.Normalize()
		if u.CheckRandomChance(u.CalculateCritilalAttackChance(target)) {
			imp.Damage.Multiply(CRITICAL_DAMAGE_FACTOR)
			imp.Damage.IsCritical = true
		}
		if imp.Duration != 0 {
			tmpImp := target.AccumulateDamageImpact(imp)
			tmpImp.Chance = 0
			temporalDamage = append(temporalDamage, tmpImp)
		} else {
			instDmg := target.ApplyDamage(imp.Damage)
			instantDamage = append(instantDamage, instDmg)
		}
	}
	return instantDamage, temporalDamage
}
