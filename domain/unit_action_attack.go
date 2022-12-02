package domain

func (u *Unit) Attack(target *Unit, damage []DamageImpact) ([]Damage, []DamageImpact) {
	wasStunned := target.State.IsStunned
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
		if wasStunned || u.CheckRandomChance(u.CalculateCritilalAttackChance(target)) {
			imp.Damage.Multiply(CRITICAL_DAMAGE_FACTOR)
			imp.Damage.IsCritical = true
		}
		imp.Chance = 0
		if imp.Duration != 0 {
			tmpImp := target.accumulateDamageImpact(imp)
			temporalDamage = append(temporalDamage, tmpImp)
		} else {
			instDmg := target.applyDamage(imp.Damage)
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

func (u *Unit) applyDamage(damage Damage) Damage {
	if damage.HasPhysicalEffect() {
		u.Inventory.IncreaseArmorWearout(true)
		u.Inventory.UpdateEquipmentByWeareout()
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
