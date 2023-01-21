package domain

import "jrpg-gang/util"

func (u *Unit) Attack(target *Unit, damage []DamageImpact) ([]Damage, []DamageImpact) {
	wasStunned := target.State.IsStunned
	instantDamage := []Damage{}
	temporalDamage := []DamageImpact{}
	totalModification := u.TotalModification()
	totalModification.Attributes.Accumulate(u.Stats.Attributes)
	totalModification.Attributes.Normalize()
	for n, imp := range damage {
		if !u.CheckRandomChance(u.CalculateAttackChance(target, imp)) {
			if n != 0 || !u.CheckRandomChance(u.CalculateCriticalMissChance()) {
				break
			}
			imp.Damage.IsCriticalMiss = true
			target = u
			wasStunned = target.State.IsStunned
		}
		if n == 0 && (wasStunned || imp.Damage.IsCriticalMiss || u.CheckRandomChance(u.CalculateCritilalAttackChance(target))) {
			imp.Damage.IsCritical = true
		}
		imp.Enchance(totalModification.Attributes, totalModification.Damage)
		imp.Normalize()
		if imp.Damage.IsCritical {
			imp.Damage.MultiplyAll(CRITICAL_DAMAGE_FACTOR)
		}
		imp.Chance = 0
		if imp.Duration != 0 {
			tmpImp := target.accumulateDamageImpact(imp)
			temporalDamage = append(temporalDamage, tmpImp)
		} else {
			instDmg := target.applyInstantDamage(imp.Damage)
			if !wasStunned && u.CheckRandomChance(target.CalculateStunChance(u, instDmg)) {
				target.State.IsStunned = true
				instDmg.WithStun = true
			}
			instantDamage = append(instantDamage, instDmg)
		}
	}
	if wasStunned && len(instantDamage) != 0 {
		target.State.IsStunned = false
	}
	return instantDamage, temporalDamage
}

func (u *Unit) applyInstantDamage(damage Damage) Damage {
	modResistance := u.TotalUnitModification().Resistance
	modResistance.Normalize()
	damage.Reduce(modResistance.Damage)
	damage.Normalize()
	if damage.HasPhysicalEffect() {
		u.Inventory.IncreaseArmorWearout(true)
		u.Inventory.UpdateEquipmentByWeareout()
	}
	resistance := u.TotalEquipmentModification().Resistance
	resistance.Accumulate(u.Stats.TotalResistance())
	resistance.Normalize()
	exhaustion := resistance.PhysicalAbsorption(damage) - modResistance.Exhaustion
	damage.Exhaustion += util.Max(exhaustion, 0)
	damage.Reduce(resistance.Damage)
	damage.Normalize()
	if damage.HasEffect() {
		damage.Apply(&u.State)
	}
	return damage
}

func (u *Unit) accumulateDamageImpact(damage DamageImpact) DamageImpact {
	resistance := u.TotalModification().Resistance
	resistance.Accumulate(u.Stats.TotalResistance())
	resistance.Normalize()
	damage.Reduce(resistance.Damage)
	damage.Normalize()
	if damage.HasEffect() {
		u.Damage = append(u.Damage, damage)
	}
	return damage
}
